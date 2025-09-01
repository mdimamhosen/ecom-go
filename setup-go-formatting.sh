#!/bin/bash

# Setup script for Go formatting and development tools
# This script installs all necessary tools for Go development with proper formatting

echo "🚀 Setting up Go development environment with formatting tools..."
echo ""

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Please install Go first."
    exit 1
fi

echo "✅ Go version: $(go version)"
echo ""

# Install gofmt (usually comes with Go)
echo "📦 Checking gofmt..."
if command -v gofmt &> /dev/null; then
    echo "✅ gofmt is available"
else
    echo "❌ gofmt not found"
fi

# Install goimports
echo "📦 Installing goimports..."
go install golang.org/x/tools/cmd/goimports@latest
if command -v goimports &> /dev/null; then
    echo "✅ goimports installed successfully"
else
    echo "❌ Failed to install goimports"
fi

# Install golangci-lint
echo "📦 Installing golangci-lint..."
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
if command -v golangci-lint &> /dev/null; then
    echo "✅ golangci-lint installed successfully"
    echo "   Version: $(golangci-lint version --format short)"
else
    echo "❌ Failed to install golangci-lint"
fi

# Install gomarkdoc for documentation
echo "📦 Installing gomarkdoc..."
go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest
if command -v gomarkdoc &> /dev/null; then
    echo "✅ gomarkdoc installed successfully"
else
    echo "❌ Failed to install gomarkdoc"
fi

# Install other useful Go tools
echo "📦 Installing additional Go tools..."

# gofumpt - stricter gofmt
go install mvdan.cc/gofumpt@latest

# gosec - security analyzer
go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest

# staticcheck - static analysis
go install honnef.co/go/tools/cmd/staticcheck@latest

echo ""
echo "🎨 Testing formatting tools..."

# Test gofmt
echo "Testing gofmt..."
if ls *.go > /dev/null 2>&1; then
    for file in *.go; do
        if [ -f "$file" ]; then
            echo "  Formatting $file with gofmt..."
            gofmt -s -w "$file"
        fi
    done
    echo "✅ gofmt test completed"
else
    echo "ℹ️  No Go files found to test formatting"
fi

# Test goimports
echo "Testing goimports..."
if ls *.go > /dev/null 2>&1; then
    for file in *.go; do
        if [ -f "$file" ]; then
            echo "  Organizing imports for $file..."
            goimports -w "$file"
        fi
    done
    echo "✅ goimports test completed"
else
    echo "ℹ️  No Go files found to test import organization"
fi

# Test golangci-lint
echo "Testing golangci-lint..."
if ls *.go > /dev/null 2>&1; then
    golangci-lint run --timeout=30s
    echo "✅ golangci-lint test completed"
else
    echo "ℹ️  No Go files found to test linting"
fi

echo ""
echo "🎯 Setup completed! Available tools:"
echo ""
echo "📋 Formatting Commands:"
echo "  gofmt -s -w .          # Format all Go files"
echo "  goimports -w .         # Organize imports"
echo "  gofumpt -w .           # Stricter formatting"
echo ""
echo "🔍 Linting Commands:"
echo "  golangci-lint run      # Comprehensive linting"
echo "  gosec ./...            # Security analysis"
echo "  staticcheck ./...      # Static analysis"
echo ""
echo "🛠️  Makefile Commands:"
echo "  make format            # Format all files"
echo "  make lint              # Run all linters"
echo "  make check             # Format + lint + test"
echo "  make all               # Complete workflow"
echo ""
echo "⚙️  VS Code Integration:"
echo "  - Auto-format on save is enabled"
echo "  - Imports organized automatically"
echo "  - Linting runs in background"
echo ""
echo "🎉 Your Go development environment is ready!"
echo ""
echo "💡 Quick start:"
echo "  1. Create a new Go file: ./new-class.sh \"My Topic\""
echo "  2. Edit your code in VS Code (auto-formatting enabled)"
echo "  3. Run: make check (format + lint + test)"
echo "  4. Generate docs: ./generate-class-doc.sh classXX.go"
echo "  5. Push to GitHub: git add . && git commit && git push"
