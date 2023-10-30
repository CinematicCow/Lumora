# Go variables
GOCMD := go
GOBUILD := $(GOCMD) build
GOTEST := $(GOCMD) test
GOCLEAN := $(GOCMD) clean
GOGET := $(GOCMD) get
GOMOD := $(GOCMD) mod
BINARY_NAME := lumora

# Test variables
TEST_DIR := ./tests
MODULE_NAME ?= all

# Main target
all: clean build test

# Build your Go application
build:
	$(GOBUILD) -o $(BINARY_NAME)

# Test the application
test:
ifeq ($(MODULE_NAME), all)
	$(GOCLEAN) -cache
	$(GOTEST) ./tests/...
else
	$(GOCLEAN) -cache
	$(GOTEST) ./tests/...$(MODULE_NAME)
endif

# Clean the workspace
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Get all dependencies
get:
	$(GOGET)

# Update go modules
update:
	$(GOMOD) tidy

init:
	./scripts/init.sh

clear:
	./scripts/clear.sh -y