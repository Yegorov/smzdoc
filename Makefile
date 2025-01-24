.DEFAULT_GOAL := all
.PHONY: run fmt vet tidy test build upgrade vendor

run:
	go run main.go

fmt:
	gofmt -w -s .

vet:
	go vet ./...

tidy:
	go mod tidy

test:
	go test -v ./...

build:
	go build -o smzdoc

upgrade:
	go get -u
	go mod tidy

vendor:
	go mod vendor

all: fmt vet tidy test build
