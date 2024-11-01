# Define the application name
APP_NAME = goDB

# Detect the OS and architecture of the host system
ifeq ($(OS),Windows_NT)
    GOOS = windows
    GOARCH = amd64
    EXT = .exe
else
    UNAME_S := $(shell uname -s)
    UNAME_M := $(shell uname -m)

    ifeq ($(UNAME_S),Linux)
        GOOS = linux
    else ifeq ($(UNAME_S),Darwin)
        GOOS = darwin
    endif

    ifeq ($(UNAME_M),x86_64)
        GOARCH = amd64
    else ifeq ($(UNAME_M),arm64)
        GOARCH = arm64
    else ifeq ($(UNAME_M),arm)
        GOARCH = arm
    else
        GOARCH = 386
    endif
    EXT =
endif

# Define the build output directory
BUILD_DIR = bin

# Default target
all: build

# Define the folders to generate tests for
generate-tests:
	gotests -w -all mongodb
	@echo "Test cases generated for mongodb"
	gotests -w -all postgres
	@echo "Test cases generated for postgres"


# Test target
test: generate-tests
	@echo "Running tests..."
	go test ./...

# Build target
build: test
	@echo "Building $(APP_NAME) for OS: $(GOOS), Architecture: $(GOARCH)"
	go build -o $(BUILD_DIR)/$(APP_NAME)$(EXT) .

# Clean target
clean:
	@echo "Cleaning build directory..."
	rm -rf $(BUILD_DIR)

# Run target
run: build
	@echo "Running $(APP_NAME)..."
	./$(BUILD_DIR)/$(APP_NAME)$(EXT)

# Cross-compile targets for other OS and architectures
build-linux: test
	set GOOS=linux 
	set GOARCH=amd64 
	go build -o $(BUILD_DIR)/$(APP_NAME)-linux .

build-darwin: test
	set GOOS=darwin 
	set GOARCH=amd64 
	go build -o $(BUILD_DIR)/$(APP_NAME)-darwin .

build-windows: test
	set GOOS=windows 
	set GOARCH=amd64 
	go build -o $(BUILD_DIR)/$(APP_NAME)-windows.exe

#build for all
build-all: build-windows build-darwin build-linux

# Help target to list all commands
help:
	@echo "Makefile commands:"
	@echo "  mingw32-make build         - Build for the local OS and architecture"
	@echo "  mingw32-make run           - Build and run the application"
	@echo "  mingw32-make clean         - Remove the build directory"
	@echo "  mingw32-make generate-tests - Generates tests"
	@echo "  mingw32-make test          - Run all tests"
	@echo "  mingw32-make build-all   - Build for Linux, Macos, Windows"
	@echo "  mingw32-make build-linux   - Build for Linux"
	@echo "  mingw32-make build-darwin  - Build for macOS"
	@echo "  mingw32-make build-windows - Build for Windows"
	@echo "  mingw32-make help          - Show this help message"