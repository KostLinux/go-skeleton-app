# Go Skeleton Application (QA)

Go Skeleton QA Application is used for the QA technical check for the necessary knowledge about automation.

Available Make functions:

- make build - Compile the code
- make run - Run the code without Hot Reloader
- make fumpt - Run gofumpt
- make mod-vendor - Run go mod vendor
- make linter - Run golangci-lint
- make gosec - Run gosec
- make validate - Run all linter, gosec and tests

## Getting started

**1. Clone the repository**

```
git clone git@github.com:KostLinux/go-skeleton-app.git -b qa
```

**2. Run the application**

Run the application via make run

```
make run
```

## QA Phase

1. Why does /not-bad endpoint fail even with header?
2. Try to look at the headers in request using Curl, what will you notice? (hint: Status)
3. Write test cases for the controller package

The docs are available [here](http://localhost:8080/swagger/index.html)