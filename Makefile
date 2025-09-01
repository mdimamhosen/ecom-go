# Makefile for Go project formatting, linting, and building
# Usage: make format, make lint, make build, etc.

.PHONY: help format lint test build clean install-tools check all

# Default target
help:
	@echo "Available commands:"
	@echo "  format       - Format all Go files using gofmt and goimports"
	@echo "  lint         - Run linting using golangci-lint"
	@echo "  test         - Run all tests"
	@echo "  build        - Build all Go files"
	@echo "  clean        - Clean build artifacts"
	@echo "  install-tools- Install required Go tools"
	@echo "  check        - Run format, lint, and test"
	@echo "  all          - Run check and build"
	@echo "  docs         - Generate documentation"

# Format Go code
format:
	@echo "🎨 Formatting Go code..."
	@if command -v goimports >/dev/null 2>&1; then \
		echo "Using goimports for formatting..."; \
		gofmt -s -w .; \
		goimports -w .; \
	else \
		echo "Using gofmt for formatting (goimports not available)..."; \
		gofmt -s -w .; \
	fi
	@echo "✅ Code formatted successfully"

# Run linting
lint:
	@echo "🔍 Running linting..."
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "Running basic go vet (golangci-lint not available)..."; \
		go vet ./...; \
	fi
	@echo "✅ Linting completed"

# Run tests
test:
	@echo "🧪 Running tests..."
	@go test -v ./...
	@echo "✅ Tests completed"

# Build all Go files
build:
	@echo "🔨 Building Go files..."
	@for file in *.go; do \
		if [ -f "$$file" ] && [ "$$file" != "class-template.go" ]; then \
			echo "Building $$file..."; \
			go build -o "$${file%.go}" "$$file" || true; \
		fi \
	done
	@echo "✅ Build completed"

# Clean build artifacts
clean:
	@echo "🧹 Cleaning build artifacts..."
	@rm -f class[0-9]* main *.exe
	@echo "✅ Clean completed"

# Install required Go tools
install-tools:
	@echo "📦 Installing Go tools..."
	@go install golang.org/x/tools/cmd/goimports@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
	@echo "✅ Tools installed"

# Check code quality (format + lint + test)
check: format lint test

# Run all checks and build
all: check build

# Generate documentation
docs:
	@echo "📚 Generating documentation..."
	@if [ -f class47.go ]; then ./generate-class-doc.sh class47.go; fi
	@if [ -f class48.go ]; then ./generate-class-doc.sh class48.go; fi
	@for file in class[0-9]*.go; do \
		if [ -f "$$file" ] && [ "$$file" != "class47.go" ] && [ "$$file" != "class48.go" ]; then \
			./generate-class-doc.sh "$$file"; \
		fi \
	done
	@echo "✅ Documentation generated"

# Format specific file
format-file:
	@if [ -z "$(FILE)" ]; then \
		echo "❌ Usage: make format-file FILE=filename.go"; \
	else \
		echo "🎨 Formatting $(FILE)..."; \
		gofmt -s -w $(FILE); \
		goimports -w $(FILE); \
		echo "✅ $(FILE) formatted"; \
	fi

# Lint specific file
lint-file:
	@if [ -z "$(FILE)" ]; then \
		echo "❌ Usage: make lint-file FILE=filename.go"; \
	else \
		echo "🔍 Linting $(FILE)..."; \
		golangci-lint run $(FILE); \
		echo "✅ $(FILE) linted"; \
	fi

# Quick development workflow
dev: format lint
	@echo "🚀 Ready for development!"
