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
set aws env variables to target your aws accounts

# run tests locally 
`go test ./tests -v`
additionally you can scope the tests with go test tags. Ex: `go test ./tests tags=unit`


what's next ?:
- s3 bucket test
- github action with tftest and tfscan.