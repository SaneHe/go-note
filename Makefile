PROJECT="server"

GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

default:
	echo ${PROJECT}; \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${PROJECT} .

fmt:
	@gofmt -s -w ${GOFILES}

install:
	@go mod tidy

test: install
	@go test ./

docker:
    @docker build -t server:latest .

.PHONY: default fmt install test build docker