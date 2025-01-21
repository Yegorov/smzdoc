.DEFAULT_GOAL := all
.PHONY: run fmt lint test build

run:
	go run main.go

fmt:
	gofmt -w -s .

lint:
	go vet ./...

test:
	go test -v ./...

build:
	go build -o smzdoc

all: fmt lint test build
