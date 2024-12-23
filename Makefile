.PHONY: build clean test run fmt vet lint

# Build variables
BINARY_NAME=ai-agentic-monitor
BUILD_DIR=build

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet
GOFMT=$(GOCMD) fmt
GOLINT=$(shell which golangci-lint 2>/dev/null || echo "$(GO_PATH)/bin/golangci-lint")
GO_PATH=$(shell go env GOPATH)

# Build the project
build:
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/main.go

# Clean build files
clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

# Run tests
test:
	$(GOTEST) -v ./...

# Run the application
run:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/main.go
	./$(BUILD_DIR)/$(BINARY_NAME)

# Format code
fmt:
	$(GOFMT) ./...

# Run go vet
vet:
	$(GOVET) ./...

# Run golangci-lint
lint:
	@if ! which golangci-lint >/dev/null; then \
		curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GO_PATH)/bin v1.62.2; \
	fi
	$(GOLINT) run ./...

# Install dependencies
deps:
	$(GOGET) -v -t -d ./...
	$(GOCMD) mod tidy

# All (format, vet, test, build)
all: fmt vet lint test build 