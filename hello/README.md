# Go Module Sample
## Setup
- Create a hello.go and implement the Hello function
- Create a hello_test.go and implement the TestHello function
- Without a go.mod file we cannot execute the test
- Create the go.mod and init it
```sh
go mod init example.com/hello
go mod tidy
```
- Run the test
```sh
go test
```

