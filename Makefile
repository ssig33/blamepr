.PHONY: build install test clean

# Binary name
BINARY_NAME=blamepr

# Build
build:
	go build -o bin/$(BINARY_NAME) ./cmd/blamepr

# Install
install:
	go install ./cmd/blamepr

# Test
test:
	go test ./...

# Clean
clean:
	rm -f bin/$(BINARY_NAME)
	go clean