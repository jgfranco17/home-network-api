project := "home-network-api"

# Default command
default:
    @just --list

# Run debug server
run-debug port:
    @echo "Starting prototype server on port {{port}}..."
    go run service/cmd/main.go --port={{port}} --dev=true

# Run production server
run-prod port:
    @echo "Starting production server on port {{port}}..."
    go run service/cmd/main.go --port {{port}} --dev=false

# Execute unit tests
test:
    @echo "Running unit tests!"
    go clean -testcache
    go test -cover github.com/jgfranco17/{{project}}/...

# Build Docker image
build:
	@echo "Building Docker image..."
	docker build -t {{project}}:latest -f ./Dockerfile .
	@echo "Docker image built successfully!"

tidy:
  -cd core && go mod tidy
  -cd service && go mod tidy
  -go mod tidy
  go work sync
