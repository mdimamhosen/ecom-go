---
title: "API Reference"
weight: 20
---

# API Reference

Complete reference for all available endpoints in the Go HTTP Server.

## Base URL

```
http://localhost:3000
```

## Endpoints

### GET /

Returns a welcome message from the server.

**Response Headers:**
- `Content-Type: application/json`

**Response:**
```
Welcome to the home page! This is a simple HTTP server in Go.
```

**Example Request:**
```bash
curl http://localhost:3000/
```

---

### GET /products

Retrieves all available products in JSON format.

**Method:** `GET`

**Response Headers:**
- `Access-Control-Allow-Origin: *`
- `Content-Type: application/json`

**Response Format:**
```json
[
  {
    "ID": 1,
    "TITLE": "Orange",
    "DESCRIPTION": "Very Sweet fruit",
    "PRICE": 100,
    "ImgUrl": "https://images.unsplash.com/photo-1557800636-894a64c1696f?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8b3JhbmdlfGVufDB8fDB8fHww"
  }
]
```

**Example Request:**
```bash
curl http://localhost:3000/products
```

**Error Responses:**

If a non-GET method is used:
- **Status Code:** `400 Bad Request`
- **Response:** `Please give me get request`

**Example Error Request:**
```bash
curl -X POST http://localhost:3000/products
```

## Data Models

### TProduct

Represents a product in the system.

| Field | Type | Description |
|-------|------|-------------|
| `ID` | `int` | Unique product identifier |
| `TITLE` | `string` | Product name/title |
| `DESCRIPTION` | `string` | Product description |
| `PRICE` | `float64` | Product price |
| `ImgUrl` | `string` | URL to product image |

**Example:**
```json
{
  "ID": 1,
  "TITLE": "Orange",
  "DESCRIPTION": "Very Sweet fruit", 
  "PRICE": 100,
  "ImgUrl": "https://images.unsplash.com/photo-1557800636-894a64c1696f?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8b3JhbmdlfGVufDB8fDB8fHww"
}
```

## Status Codes

| Code | Description |
|------|-------------|
| `200` | Success - Request completed successfully |
| `400` | Bad Request - Invalid request method |

## CORS Support

The API supports Cross-Origin Resource Sharing (CORS) with the following header:
```
Access-Control-Allow-Origin: *
```

This allows the API to be called from any domain, making it suitable for frontend integration.