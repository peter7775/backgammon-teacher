# Makefile for SQL Graph Visualizer
# Requires Go 1.24+

.PHONY: help format sec-scan ci-check lint

# Default target
help:
	@echo "Backgammon Teacher - Development Commands"
	@echo ""
	@echo "Available targets:"
	@echo "  format     - Format Go code"
	@echo "  lint       - Run golangci-lint"
	@echo "  build      - Build the application"
	@echo "  sec-scan   - Run security scans (govulncheck, gosec)"
	@echo "  ci-check   - Run CI checks locally"
	@echo ""

format:
	@echo "Formatting Go code..."
	gofmt -s -w .
	@echo "Code formatted"

# Build the unified CLI application
build:
	@echo "Building Backgammon teacher..."
	go build -o backgammon-teacher ./...
	@echo "Build completed: ./backgammon-teacher"


# Run golangci-lint
lint:
	@echo "Running golangci-lint..."
	$(HOME)/go/bin/golangci-lint run --timeout=10m
	@echo "Lint completed"

# Run CI checks locally
ci-check: install generate format
	@echo "Running CI checks locally..."
	@echo "Checking Go modules consistency..."
	go mod tidy
	@if [ -n "$$(git diff go.mod go.sum)" ]; then \
		echo "go.mod or go.sum are not up to date"; \
		git diff go.mod go.sum; \
		exit 1; \
	fi
	@echo "Checking code formatting..."
	@if [ -n "$$(gofmt -s -l .)" ]; then \
		echo "Code is not formatted properly:"; \
		gofmt -s -d .; \
		exit 1; \
	fi
	@echo "Running go vet..."
	go vet ./...
	@echo "Running golangci-lint..."
	$(HOME)/go/bin/golangci-lint run --timeout=10m
	@echo "Building..."
	go build -v ./...
	@echo "Running tests..."
	$(MAKE) test
	@echo "All CI checks passed"
	@echo "environment ready"

# Run security scans
sec-scan:
	@echo "🔍 Running security scans..."
	@echo "Installing security tools..."
	go clean -cache -modcache -testcache
	go install golang.org/x/vuln/cmd/govulncheck@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	@echo "Running govulncheck..."
	$(HOME)/go/bin/govulncheck ./...
	@echo "Running gosec..."
	$(HOME)/go/bin/gosec -exclude=G104,G115,G201,G204,G301,G304,G306 -exclude-dir=internal/interfaces/graphql/generated ./...
	@echo "✅ Security scans completed"

