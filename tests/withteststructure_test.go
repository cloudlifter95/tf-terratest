package test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/aws"
	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
	"github.com/stretchr/testify/assert"
)

func TestS3(t *testing.T) {
	// now := time.Now()
	expectedName := "mjtestbucket1234567ashq8123"

	expectedEnvironment := "Dev"

	awsRegion := "eu-central-1"
	awsProfile := "moj9-pg"
	os.Setenv("AWS_PROFILE", awsProfile)

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../terraform/",

		Vars: map[string]interface{}{
			"tag_bucket_name":        expectedName,
			"tag_bucket_environment": expectedEnvironment,
		},
	})

	defer terraform.Destroy(t, terraformOptions)

	terraform.InitAndApply(t, terraformOptions)

	// Check versioning
	test_structure.RunTestStage(t, "test_versioning", func() {
		bucketID := terraform.Output(t, terraformOptions, "bucket_id")
		bucketVersionValidation(t, terraformOptions, awsRegion, bucketID)
	})

	// Tags comparison
	test_structure.RunTestStage(t, "test_tags", func() {
		tagsValidation(t, terraformOptions)
	})
}

func bucketVersionValidation(t *testing.T, terraformOptions *terraform.Options, awsRegion string, bucketID string) {
	// Bucket versioning comparison
	actualStatus := aws.GetS3BucketVersioning(t, awsRegion, bucketID)
	expectedStatus := "Enabled"
	assert.Equal(t, expectedStatus, actualStatus)
}

func tagsValidation(t *testing.T, terraformOptions *terraform.Options) {
	tagsMap := terraform.OutputMap(t, terraformOptions, "tags")

	tagsToCheck := []string{"Environment", "Name"}

	filteredTags := make(map[string]string)

	for _, tag := range tagsToCheck {
		if value, ok := tagsMap[tag]; ok {
			filteredTags[tag] = value
		}
	}

	expectedTagsString := `{"Environment":"Dev","Name":"mjtestbucket1234567ashq8123"}`

	var expectedTags map[string]string
	err := json.Unmarshal([]byte(expectedTagsString), &expectedTags)
	if err != nil {
		t.Fatalf("Failed to unmarshal expected tags: %s", err)
	}

	assert.Equal(t, expectedTags, filteredTags)
}
