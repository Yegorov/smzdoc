.DEFAULT_GOAL := all
.PHONY: run fmt vet tidy test build

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

all: fmt vet tidy test build
