# Go HTTP Server - Line-by-Line Code Explanation

This document provides a detailed line-by-line explanation of `class47.go`, a simple HTTP server written in Go.

## 📦 Package and Imports

```go
package main
```
- Declares this as the main package, making it an executable program

```go
import (
    "encoding/json"    // For JSON encoding/decoding
    "fmt"             // For formatted I/O operations  
    "net/http"        // For HTTP server functionality
)
```

## 🏠 Home Handler Function

```go
func homeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Welcome to the home page! This is a simple HTTP server in Go.")
}
```

**Line-by-line breakdown:**
- `func homeHandler(w http.ResponseWriter, r *http.Request)` - Function signature for HTTP handler
- `w.Header().Set("Content-Type", "application/json")` - Sets response content type to JSON
- `w.WriteHeader(http.StatusOK)` - Sets HTTP status code to 200 (OK)
- `fmt.Fprintln(w, "...")` - Writes welcome message to response body

## 🏷️ Product Type Definition

```go
type TProduct struct {
    ID          int
    TITLE       string
    DESCRIPTION string
    PRICE       float64
    ImgUrl      string
}
```

**Explanation:**
- Defines a struct type `TProduct` with 5 fields
- `ID` - Integer for unique product identification
- `TITLE` - String for product name
- `DESCRIPTION` - String for product details
- `PRICE` - Float64 for product price
- `ImgUrl` - String for product image URL

## 📦 Products Handler Function

```go
func getProducts(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Controll-Allow-Origin", "*")     // CORS header (typo: should be "Control")
    w.Header().Set("Content-Type", "application/json")      // JSON content type
    
    if r.Method != http.MethodGet {                         // Check if request method is GET
        http.Error(w, "Please give me get request", 400)    // Return 400 error for non-GET requests
        return
    }

    w.WriteHeader(http.StatusOK)    // Set 200 OK status
    encoder := json.NewEncoder(w)   // Create JSON encoder that writes to response
    encoder.Encode(productList)     // Encode productList slice to JSON and send
}
```

## 🗃️ Global Variable

```go
var productList []TProduct
```
- Declares a global slice to store TProduct structs
- Initially empty, populated in `init()` function

## 🎯 Main Function - Detailed Breakdown

```go
func main() {
```

### Step 1: Create Router
```go
mux := http.NewServeMux() // router
```
- `http.NewServeMux()` - Creates a new HTTP request multiplexer (router)
- `mux` - Variable that holds the router instance
- Router is responsible for matching incoming requests to appropriate handlers

### Step 2: Register Route Handlers
```go
mux.HandleFunc("/", homeHandler) // home route
```
- `mux.HandleFunc("/", homeHandler)` - Registers `homeHandler` function for root path "/"
- When someone visits `http://localhost:3000/`, `homeHandler` will be called

```go
mux.HandleFunc("/products", getProducts)
```
- Registers `getProducts` function for "/products" path
- When someone visits `http://localhost:3000/products`, `getProducts` will be called

### Step 3: Server Startup Message
```go
fmt.Println("Server is running on port 3000")
```
- Prints a message to console indicating server is starting on port 3000

### Step 4: Start HTTP Server
```go
err := http.ListenAndServe(":3000", mux)
```
- `http.ListenAndServe(":3000", mux)` - Starts HTTP server on port 3000
- `:3000` - Tells server to listen on port 3000 (localhost:3000)
- `mux` - Uses our router to handle incoming requests
- Returns an error if server fails to start
- This is a **blocking call** - program waits here and handles requests

### Step 5: Error Handling
```go
if err != nil {
    fmt.Println("Error starting server:", err)
    return
}
```
- Checks if server startup returned an error
- If error exists, prints error message and exits main function
- `return` - Exits the main function, ending the program

## 🔧 Initialization Function

```go
func init() {
    prod1 := TProduct{
        ID:          1,
        TITLE:       "Orange",
        DESCRIPTION: "Very Sweet fruit",
        PRICE:       100,
        ImgUrl:      "https://images.unsplash.com/photo-1557800636-894a64c1696f?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8b3JhbmdlfGVufDB8fDB8fHww",
    }

    productList = append(productList, prod1)
}
```

**Explanation:**
- `func init()` - Special function that runs automatically before `main()`
- Creates a sample product with ID=1, title="Orange", etc.
- `productList = append(productList, prod1)` - Adds the product to global productList slice

## � Program Flow

1. **Program starts** → `init()` runs first, populating productList
2. **main() begins** → Creates router and registers handlers  
3. **Server starts** → Listens on port 3000 for incoming requests
4. **Request handling** → Router directs requests to appropriate handlers
5. **Response** → Handlers process requests and send responses back

## 🏃‍♂️ How to Run

### Prerequisites
- Go installed on your system

### Steps to Run:
1. Navigate to the project directory
2. Run the command:
   ```bash
   go run class47.go
   ```
3. You'll see: "Server is running on port 3000"
4. Open browser or use curl to test

### Testing the Endpoints:

**Test Home Page:**
```bash
curl http://localhost:3000/
```
Response: "Welcome to the home page! This is a simple HTTP server in Go."

**Test Products API:**
```bash
curl http://localhost:3000/products
```
Response: JSON array with the Orange product

**Test Error Handling (POST request to products):**
```bash
curl -X POST http://localhost:3000/products
```
Response: "Please give me get request" with 400 status

## 📝 Key Learning Points

1. **HTTP Server Setup**: Using `http.NewServeMux()` as router
2. **Handler Functions**: Functions that handle HTTP requests
3. **JSON Responses**: Using `json.NewEncoder()` to send JSON
4. **Method Validation**: Checking request methods before processing
5. **CORS Headers**: Allowing cross-origin requests
6. **Global Variables**: Sharing data between functions
7. **Init Function**: Code that runs before main()

## 🔍 Important Notes

- The server runs on port 3000
- Only GET requests are allowed for `/products`
- Product data is stored in memory (resets when server restarts)
- There's a typo in CORS header: "Access-Controll-Allow-Origin" should be "Access-Control-Allow-Origin"

---

This code demonstrates basic HTTP server concepts in Go and is perfect for learning web development fundamentals.
