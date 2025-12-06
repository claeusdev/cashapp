.PHONY: swagger build docker-build docker-up docker-down

# Generate Swagger documentation
swagger:
	@echo "Generating Swagger documentation..."
	@swag init -g main.go -o ./docs

# Build the application
build:
	@echo "Building application..."
	@go build -o main .

# Build Docker image
docker-build:
	@echo "Building Docker image..."
	@docker-compose build

# Start Docker containers
docker-up:
	@echo "Starting Docker containers..."
	@docker-compose up -d

# Stop Docker containers
docker-down:
	@echo "Stopping Docker containers..."
	@docker-compose down

# Run the application locally
run:
	@echo "Running application..."
	@go run main.go

# Install dependencies
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Install swagger CLI (if not installed)
install-swagger:
	@echo "Installing swagger CLI..."
	@go install github.com/swaggo/swag/cmd/swag@latest

