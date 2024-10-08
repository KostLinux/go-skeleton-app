# Go Skeleton Application (QA)

Go Skeleton QA Application is used for the QA technical check for the necessary knowledge about automation.

The application is built using basic MVC layered architecture.

Available Make functions:

- make build - Compile the code
- make run - Run the code without Hot Reloader
- make fumpt - Run gofumpt
- make mod-vendor - Run go mod vendor
- make linter - Run golangci-lint
- make gosec - Run gosec
- make validate - Run all linter, gosec and tests

## Prerequisites

- [Go](https://go.dev/dl/) installed

### Installing Go

Golang installation guide can be seen in [official docs](https://go.dev/doc/install)

The installation files can be found [under downloads](https://go.dev/dl/)

```
// MacOS
brew install go

// Linux
On Ubuntu/Debian-based distros: `sudo apt install golang`
On RHEL/Fedora: `sudo dnf install golang`
You can also use Snap to install Go: `sudo snap install go --classic`
```

## Getting started

**1. Clone the repository**

```
// Via SSH
git clone git@github.com:KostLinux/go-skeleton-app.git -b qa

// Via HTTPS
git clone https://github.com/KostLinux/go-skeleton-app.git -b qa
```

**2. Run the application**

Run the application via `go run main.go`

```
go run main.go
```

## QA Phase

1. Why does /notbad endpoint fail even with header?

2. Try to look at the headers in response using Curl, what will you notice? 

Make a request `curl http://localhost:8080/ -I` & `http://localhost:8080/` and compare the responses.

Find out, why HTTP status codes are different.

3. Write a test case, which tests whether the rate limiter works properly

Write a simple non-functional test using your preferred language to check whether the rate limiter works.

4. Write test cases for the controller package (optional)

The docs are available [here](http://localhost:8080/swagger/index.html)