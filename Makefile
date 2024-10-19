# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOTEST = $(GOCMD) test
GORUN = $(GOCMD) run

# Binary name
BINARY_NAME = spb

# Build target
build:
	$(GOBUILD) -o ./bin/$(BINARY_NAME) ./cmd

# Run target
run: build
	./bin/$(BINARY_NAME)

# Test target
test:
	$(GOTEST) ./...

# Clean target
clean:
	rm -f ./bin/$(BINARY_NAME)

.PHONY: build run test clean
