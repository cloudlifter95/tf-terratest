WIP
# terratest - tfsec repository
This repo is intended to try and demonstrate some of the capabilities of test tools built on top of terraform. We intend to try the following tools:
- static testing: **tfsec**
- dynamic testing: **terratest**
This project comes with an automated github runner, as well as a manual section with command lines.


# install dependencies
from project root folder, run: 
`go mod download`  

# credentials
- remove profile from providers.tf
- set aws env variables to target your aws accounts

# run tests locally 
`go test ./tests -v`
additionally you can scope the tests with go test tags. Ex: `go test ./tests tags=unit`
(delete cache if tags separation is not working: `go clean -modcache`)

# terratest stages - test_structure
terratest test_structure enables the sequencing of tests without provisioning and deprovisioning the infrastructure.

For that a main test function is defined, which:
- provision and destroys infra
- implements test_structure to define the sequences of tests to be performed.

skip a test with: `SKIP_<stagename>=1 go test ./tests/ -v`
EX: `SKIP_test_versioning=1 go test ./tests/ -v`