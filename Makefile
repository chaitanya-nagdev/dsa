# Define variables
BINARY_NAME=dsa.out
PKG=./...

# Build the Go binary
build:
	@echo "Building the binary..."
	@go build -o $(BINARY_NAME) main.go

# Run the Go binary
run: build
	@echo "Running the application..."
	@./$(BINARY_NAME)

# Run tests in verbose mode
verbose:
	@echo "Running tests in verbose mode..."
	@go test -v $(PKG)

# Run tests with code coverage
coverage:
	@echo "Running tests with code coverage..."
	@go test -cover -coverprofile=coverage.out $(PKG)
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"
	@echo "Opening in browser for linux"
	@xdg-open coverage.html ||  @echo "Failed to open in Firefox. Please open 'coverage.html' manually." # Linux Version


# Clean up build artifacts
clean:
	@echo "Cleaning up..."
	@go clean
	@rm -f $(BINARY_NAME)

# Default target
all: build verbose 
	@echo "Build and tests completed successfully."
