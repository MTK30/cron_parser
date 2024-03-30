## Cron Expression Parser

This project parses a cron string and expands each field

## Installation
install golang with version > 1.17
<!-- in mac os  -->
    brew update&& brew install golang
    go version

<!-- in linux -->
    sudo apt install golang-go


Check installation with 
    go version

## Running of Program 
1. clone to local system and move to the directory location with cd `<yourPath>/cron_parser`
2. install any depdancy of the project with `go get` cmd
3. run the program with go run cmd `go run main.go "$cronExpressionString"`
or 
build the program with `go build` cmd
now run the exectuable 
`./cron_parser "$cronExpressionString"`


## Project Structe 
main.go -> contians basic drive/handler code
go.mod -> file listing dependancies
utils package -> common functional and utilites used acrros the project
parser package -> file wher the actual parsing of expression takes place
test -> test file to test smaller snipts of code 

## Test cases 

Individual level method based on entire flow test cases are written in `test` floder

Please check `TestParseCronExpression`  test case which tests the entire parser as a whole. 






