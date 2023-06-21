.PHONY: all test clean

test:
	go test .internal/server/ -v

lint:
	golangci-lint run

run:
	go run ./cmd/server/main.go

substring:
	go build -o substring ./cmd/cli/

build:
	docker build --pull --rm -f "Dockerfile" -t substringfinder:latest "."

up:
	docker run -p 8080:8080 substringfinder:latest