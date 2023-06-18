.PHONY: all test clean

test:
	go test .internal/server/ -v

lint:
	golangci-lint run

run:
	go run ./cmd/shop-crud-project/main.go

substring:
	go build -o substring ./cmd/cli/

build:
	docker build -t server:latest

up:
	docker run -p 8080:8080 server:latest