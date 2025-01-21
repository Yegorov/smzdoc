.DEFAULT_GOAL := all
.PHONY: run build fmt test

run:
	go run main.go

build:
	go build -o smzdoc

fmt:
	gofmt -w -s .

test:
	go test -v ./...

all: fmt test build
