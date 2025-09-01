---
title: "Introduction"
weight: 1
date: '2025-09-01T23:41:41+06:00'
draft: false
---

# Introduction to Go HTTP Server

## What is this project?

This is a simple HTTP server built in Go that demonstrates fundamental web development concepts. The server provides a RESTful API for managing product data with JSON responses.

## Key Features

### 🌐 HTTP Server
- Built with Go's standard `net/http` package
- Custom router using `http.NewServeMux()`
- Runs on port 3000

### 📡 API Endpoints
- **Home endpoint** (`/`) - Welcome message
- **Products endpoint** (`/products`) - Product data in JSON format

### 🔧 Technical Implementation
- **Type-safe structures** - Custom `TProduct` struct
- **JSON encoding** - Automatic JSON marshaling
- **Method validation** - Only GET requests allowed for products
- **CORS headers** - Cross-origin resource sharing support
- **Error handling** - Proper HTTP status codes

## Code Structure

The entire server is implemented in a single file (`class47.go`) with:
- Handler functions for each endpoint
- Product type definition
- Global product storage
- Server initialization and startup

## Learning Objectives

This project helps you understand:
- HTTP server basics in Go
- Request routing and handling
- JSON API development
- CORS configuration
- Error handling patterns
- Go's type system and structs

## Next Steps

Continue reading to explore the [detailed architecture](../architecture) and [API documentation](../../api).
