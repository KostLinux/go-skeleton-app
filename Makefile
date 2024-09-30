_DEFAULT_GOAL := run

run:
	go run main.go

build:
	go build -o bin/main main.go


fumpt:
	gofumpt -w .

mod-vendor:
	go mod vendor

linter:
	@golangci-lint run

gosec:
	@gosec -quiet ./...

validate: linter gosec