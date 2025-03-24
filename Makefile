# Makefile

# Variables
BINARY_NAME := kpop-cli
VERSION := 0.1.0

# Run app
run:
	go run cmd/kpop/main.go 12345

# Install dependencies
deps:
	go mod tidy

# Build binaries for all platforms
build:
	GOOS=linux GOARCH=amd64 go build -o dist/$(BINARY_NAME)-linux-amd64 ./cmd/kpop
	GOOS=linux GOARCH=arm64 go build -o dist/$(BINARY_NAME)-linux-arm64 ./cmd/kpop
	GOOS=darwin GOARCH=amd64 go build -o dist/$(BINARY_NAME)-darwin-amd64 ./cmd/kpop
	GOOS=darwin GOARCH=arm64 go build -o dist/$(BINARY_NAME)-darwin-arm64 ./cmd/kpop
	GOOS=windows GOARCH=amd64 go build -o dist/$(BINARY_NAME)-windows-amd64.exe ./cmd/kpop
	GOOS=windows GOARCH=arm64 go build -o dist/$(BINARY_NAME)-windows-arm64.exe ./cmd/kpop

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
