# Go variables
GOCMD := go
GOBUILD := $(GOCMD) build
GORUN := $(GOCMD) run
GOTEST := $(GOCMD) test
GOCLEAN := $(GOCMD) clean
GOGET := $(GOCMD) get
GOMOD := $(GOCMD) mod
BINARY_NAME := lumora
MAIN_DIR = ./cmd/lumora

# Test variables
TEST_DIR := ./tests
MODULE_NAME ?= all

.PHONY: tmp

# Main target
all: clean build test

# Build your Go application
build:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_DIR)

# Run you Go application 
run:
	$(GORUN) $(MAIN_DIR)


# Test the application
test:
ifeq ($(MODULE_NAME), all)
	$(GOTEST) ./tests/...
else
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

tmp:
	rm -r ./tmp/*