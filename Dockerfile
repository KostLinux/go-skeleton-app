FROM golang:1.23.0-bookworm AS builder

## Labels
LABEL "Team" "LookinLabs"

LABEL "Product" "GoApp"

WORKDIR /app

COPY . .

RUN go mod download \
    && CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o ./skeleton -a -ldflags "-s -w" -installsuffix cgo .

FROM alpine AS application

WORKDIR /app

ENV TLS_MODE="true"

COPY ./certs/ ./certs/
COPY --from=builder /app/skeleton .

EXPOSE 8080

ENTRYPOINT ["./skeleton"]