# Makefile for the geminik8s project

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool
BINARY_NAME=gemin_k8s
BINARY_PATH=./bin/$(BINARY_NAME)
MAIN_PATH=./cmd/gemin_k8s/main.go

# Build-time version variables
VERSION ?= $(shell git describe --tags --always --dirty)
COMMIT ?= $(shell git rev-parse HEAD)
DATE ?= $(shell date -u +'%Y-%m-%dT%H:%M:%SZ')
BUILT_BY ?= makefile

# LDFLAGS allows us to inject variables at build time
LDFLAGS = -ldflags "-X 'github.com/turtacn/geminik8s/internal/app/cli.Version=$(VERSION)' \
-X 'github.com/turtacn/geminik8s/internal/app/cli.Commit=$(COMMIT)' \
-X 'github.com/turtacn/geminik8s/internal/app/cli.Date=$(DATE)' \
-X 'github.com/turtacn/geminik8s/internal/app/cli.BuiltBy=$(BUILT_BY)'"

# Default target
all: build

# Build the binary
build:
	@echo "Building $(BINARY_NAME) version $(VERSION)..."
	@mkdir -p ./bin
	$(GOBUILD) $(LDFLAGS) -o $(BINARY_PATH) $(MAIN_PATH)
	@echo "$(BINARY_NAME) built successfully at $(BINARY_PATH)"

# Run tests
test:
	@echo "Running tests..."
	$(GOTEST) -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	$(GOCLEAN)
	@rm -rf ./bin
	@rm -rf ./dist
	@echo "Cleaned."

# Run the application with default arguments
run: build
	@echo "Running $(BINARY_NAME)..."
	$(BINARY_PATH) --help

# Cross-compilation for Linux
build-linux:
	@echo "Building for Linux..."
	@mkdir -p ./dist/linux
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) $(LDFLAGS) -o ./dist/linux/$(BINARY_NAME) $(MAIN_PATH)
	@echo "Linux binary at ./dist/linux/$(BINARY_NAME)"

# Tidy dependencies
deps:
	@echo "Tidying dependencies..."
	$(GOCMD) mod tidy

# Tooling
GOLANGCILINT_VERSION=v1.55.2
GOFUMPT_VERSION=v0.6.0

## install-lint: installs golangci-lint
install-lint:
	@echo "Installing golangci-lint..."
	@GOBIN=$(shell pwd)/bin $(GOCMD) install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCILINT_VERSION)

## install-fumpt: installs gofumpt
install-fumpt:
	@echo "Installing gofumpt..."
	@GOBIN=$(shell pwd)/bin $(GOCMD) install mvdan.cc/gofumpt@$(GOFUMPT_VERSION)

## tools: installs all tools
tools: install-lint install-fumpt

## lint: runs the linter
lint:
	@echo "Linting code..."
	@./bin/golangci-lint run ./...

## format: formats the code
format:
	@echo "Formatting code..."
	@./bin/gofumpt -l -w .

# Update phony targets
.PHONY: all build test clean run build-linux deps tools install-lint install-fumpt lint format

#Personal.AI order the ending
