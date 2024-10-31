# Makefile

# Variables
BINARY_NAME := kpop-cli
VERSION := 1.0.0

# Install dependencies
deps:
	go mod tidy

# Build binaries for all platforms
build:
	GOOS=linux GOARCH=amd64 go build -o dist/$(BINARY_NAME)-linux-amd64 main.go
	GOOS=darwin GOARCH=amd64 go build -o dist/$(BINARY_NAME)-darwin-amd64 main.go
	GOOS=windows GOARCH=amd64 go build -o dist/$(BINARY_NAME)-windows-amd64.exe main.go

# Run tests
test:
	go test ./...

# Clean up build artifacts
clean:
	rm -rf dist/

# Generate autocompletion scripts
completion:
	go run main.go completion bash > /etc/bash_completion.d/$(BINARY_NAME)
	go run main.go completion zsh > "${fpath[1]}/_$(BINARY_NAME)"

# Run linter
lint:
	golangci-lint run

# Install binary to local /usr/local/bin
install: build
	sudo cp dist/$(BINARY_NAME)-$(shell uname -s | tr '[:upper:]' '[:lower:]')-amd64 /usr/local/bin/$(BINARY_NAME)
