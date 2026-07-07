.PHONY: build test lint install clean

build:
	go build -o bin/exact-online-accountancy-pp-cli ./cmd/exact-online-accountancy-pp-cli

test:
	go test ./...

lint:
	golangci-lint run

install:
	go install ./cmd/exact-online-accountancy-pp-cli

clean:
	rm -rf bin/

build-mcp:
	go build -o bin/exact-online-accountancy-pp-mcp ./cmd/exact-online-accountancy-pp-mcp

install-mcp:
	go install ./cmd/exact-online-accountancy-pp-mcp

build-all: build build-mcp
