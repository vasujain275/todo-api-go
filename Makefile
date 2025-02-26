# Define environment variables
export $(shell cat .env | xargs)

# Go-related variables
GO_CMD=go
BUILD_DIR=bin
BIN_NAME=main
CMD_DIR=./cmd/api
MIGRATIONS_DIR=./cmd/migrate/migrations
DB_ADDR=$(DB_ADDR)

# Tools
MIGRATE_CMD=migrate

.PHONY: all build run clean migrate test lint fmt vet docker-build docker-run

# Default target
all: build

# Build the Go application
build:
	@echo "Building the application..."
	$(GO_CMD) build -o $(BUILD_DIR)/$(BIN_NAME) $(CMD_DIR)

# Run the application
run: build
	@echo "Starting the application..."
	./$(BUILD_DIR)/$(BIN_NAME)

# Clean build artifacts
clean:
	@echo "Cleaning up..."
	rm -rf $(BUILD_DIR)

# Run database migrations
migrate:
	@echo "Running database migrations..."
	$(MIGRATE_CMD) -path=$(MIGRATIONS_DIR) -database=$(DB_ADDR) up

# Run tests
test:
	@echo "Running tests..."
	$(GO_CMD) test ./... -cover

# Linting
lint:
	@echo "Running linter..."
	golangci-lint run

# Format the code
fmt:
	@echo "Formatting code..."
	$(GO_CMD) fmt ./...

# Run vet
vet:
	@echo "Running go vet..."
	$(GO_CMD) vet ./...

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	docker build -t my-go-app .

# Run Docker container
docker-run:
	@echo "Running Docker container..."
	docker compose up
