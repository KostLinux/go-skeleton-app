# Build image

FROM golang:1.19.13-alpine3.18 as builder

## API Port
ENV GOPROXY=http://proxy.golang.org,direct

## Labels
LABEL "Team" "LookinLabs"

LABEL "Product" "GoApp"

## Set the Current Working Directory inside the container
WORKDIR /app

## Copy the source from the current directory to the Working Directory inside the container
COPY . .

## Install git.
## Git is required for fetching the dependencies.
RUN apk add \
    --no-cache \
    --allow-untrusted \
    --repository http://dl-cdn.alpinelinux.org/alpine/v3.19/main \
    --update git bash build-base

## Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

######################
# Application image
FROM golang:1.19.13-alpine3.18

WORKDIR /app

COPY --from=builder ./app ./

EXPOSE 8080

CMD [ "go", "run", "main.go" ]