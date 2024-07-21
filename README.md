# Go Skeleton Application

Go skeleton application repository used for production-ready golang applications via golang best practices meshed with MVC architecture.

Available Make functions:

- make build - Compile the code
- make air - Run the code with Golang Hot Reloader (Air)
- make run - Run the code without Hot Reloader
- make fumpt - Run gofumpt
- make mod-vendor - Run go mod vendor
- make linter - Run golangci-lint
- make gosec - Run gosec
- make test - Run tests
- make validate - Run all linter, gosec and tests
- make migrate-up - Run database migrations
- make migrate-down - Rollback database migrations
- make migrate-create - Create a new database migration

## Getting started

You can pull the example application from [AWS ECR Repository](https://gallery.ecr.aws/c2w5h6c4/go-skeleton-app).

It doesn't use any versioning system for purpose of simplicity. You can use the following command to pull the image:

```bash
docker pull public.ecr.aws/c2w5h6c4/go-skeleton-app:latest
```

```bash
curl https://localhost:8080/ -k -v
```

The <tag> part is the version of the image you want to pull.
