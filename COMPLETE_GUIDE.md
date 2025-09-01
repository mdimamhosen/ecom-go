# Complete Guide: File Naming & Automated Documentation Workflow

## 📁 File Naming Conventions

### Go Source Files

**Pattern**: `classXX.go` where XX is a sequential number

**Examples**:
- ✅ `class47.go` - HTTP Server (your existing file)
- ✅ `class48.go` - User Management API (example created)
- ✅ `class49.go` - Authentication System
- ✅ `class50.go` - Database Integration
- ✅ `class100.go` - Advanced Features

**Important Rules**:
- Always use lowercase `class`
- Use sequential numbers (47, 48, 49, etc.)
- Always end with `.go` extension
- No spaces or special characters
- Keep it simple and consistent

### Documentation Files (Auto-Generated)

The system automatically creates:
- `docs/content/docs/classes/class47.md`
- `docs/content/docs/classes/class48.md`
- `docs/content/docs/classes/_index.md` (classes listing)

## 📝 Go File Structure Template

### Basic Template for New Classes

```go
package main

import (
    // Add your imports here
    "fmt"
    "net/http"
    "encoding/json"
)

// [ClassDescription] - Brief description of what this class demonstrates
// Example: HTTPServer demonstrates basic HTTP server functionality

// Define your types/structs here
type YourStruct struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

// Define your functions here
func YourFunction() {
    fmt.Println("Your implementation")
}

// Main function for this class (rename to avoid conflicts)
func ClassXXMain() {
    // Your main logic here
    fmt.Println("Class XX demonstration")
}

// Optional: Demo function to showcase the class
func ClassXXDemo() {
    ClassXXMain()
}
```

### Example: Complete Class49 Template

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "time"
)

// Class49 demonstrates JWT Authentication and Middleware
// This class shows how to implement authentication in Go HTTP servers

// AuthToken represents an authentication token
type AuthToken struct {
    Token     string    `json:"token"`
    ExpiresAt time.Time `json:"expires_at"`
    UserID    int       `json:"user_id"`
}

// AuthService handles authentication operations
type AuthService struct {
    tokens map[string]AuthToken
}

// NewAuthService creates a new authentication service
func NewAuthService() *AuthService {
    return &AuthService{
        tokens: make(map[string]AuthToken),
    }
}

// GenerateToken creates a new authentication token
func (as *AuthService) GenerateToken(userID int) AuthToken {
    token := AuthToken{
        Token:     fmt.Sprintf("token_%d_%d", userID, time.Now().Unix()),
        ExpiresAt: time.Now().Add(24 * time.Hour),
        UserID:    userID,
    }
    as.tokens[token.Token] = token
    return token
}

// ValidateToken checks if a token is valid
func (as *AuthService) ValidateToken(tokenString string) (AuthToken, bool) {
    token, exists := as.tokens[tokenString]
    if !exists || time.Now().After(token.ExpiresAt) {
        return AuthToken{}, false
    }
    return token, true
}

// loginHandler handles user login
func (as *AuthService) loginHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // In a real app, you'd validate credentials here
    // For demo, we'll just generate a token for user ID 1
    token := as.GenerateToken(1)
    
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Login successful",
        "token":   token,
    })
}

// protectedHandler demonstrates a protected endpoint
func (as *AuthService) protectedHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    // Check for authorization header
    authHeader := r.Header.Get("Authorization")
    if authHeader == "" {
        http.Error(w, "Authorization header required", http.StatusUnauthorized)
        return
    }

    // Validate token (remove "Bearer " prefix if present)
    tokenString := authHeader
    if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
        tokenString = authHeader[7:]
    }

    token, valid := as.ValidateToken(tokenString)
    if !valid {
        http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
        return
    }

    response := map[string]interface{}{
        "message": "Access granted to protected resource",
        "user_id": token.UserID,
        "data":    "This is protected data",
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

// Class49Main demonstrates the authentication server
func Class49Main() {
    authService := NewAuthService()

    mux := http.NewServeMux()
    mux.HandleFunc("/login", authService.loginHandler)
    mux.HandleFunc("/protected", authService.protectedHandler)
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(map[string]string{
            "message":   "Class 49 - Authentication API",
            "endpoints": "/login (POST), /protected (GET with token)",
        })
    })

    port := ":8080"
    fmt.Printf("🔐 Class 49 Auth Server starting on port %s\n", port)
    fmt.Println("📋 Endpoints:")
    fmt.Println("   POST /login      - Get authentication token")
    fmt.Println("   GET  /protected  - Access protected resource (requires token)")
    
    http.ListenAndServe(port, mux)
}
```

## 📋 Step-by-Step Workflow

### 1. Create Your New Go File

```bash
# Navigate to your project directory
cd /home/mdimamhosen/Code/GoLang/GoLangWithHabibVai/Backend

# Create your new class file (replace XX with next number)
nano classXX.go
```

### 2. File Content Guidelines

**Required Elements**:
- Package declaration: `package main`
- Descriptive comments for types and functions
- Unique function names (avoid conflicts with existing files)
- Main function renamed to `ClassXXMain()` where XX is your class number

**Recommended Structure**:
```go
// File header comment explaining what this class demonstrates
package main

// Imports section
import (...)

// Type definitions with comments
type YourType struct {...}

// Service/Helper structs if needed
type YourService struct {...}

// Constructor functions
func NewYourService() *YourService {...}

// Business logic functions with comments
func (ys *YourService) YourMethod() {...}

// HTTP handlers if it's a web service
func (ys *YourService) yourHandler(w http.ResponseWriter, r *http.Request) {...}

// Main demonstration function
func ClassXXMain() {...}
```

### 3. Generate Documentation

```bash
# Generate documentation for your new file
./generate-class-doc.sh classXX.go
```

**What this does**:
- Creates `docs/content/docs/classes/classXX.md`
- Updates the classes index page
- Extracts functions and types automatically
- Adds file metadata (size, lines, creation date)

### 4. Commit and Push

```bash
# Add all changes
git add .

# Commit with descriptive message
git commit -m "Add classXX.go: [Brief description of what it demonstrates]"

# Examples:
# git commit -m "Add class49.go: JWT Authentication and Middleware"
# git commit -m "Add class50.go: Database Integration with PostgreSQL"

# Push to GitHub
git push origin main
```

## 🎯 File Naming Examples for Different Topics

### Web Development
- `class47.go` - Basic HTTP Server ✅
- `class48.go` - User Management API ✅
- `class49.go` - JWT Authentication
- `class50.go` - Middleware Implementation
- `class51.go` - File Upload Handler
- `class52.go` - WebSocket Server

### Database
- `class60.go` - Database Connection
- `class61.go` - CRUD Operations
- `class62.go` - Database Migrations
- `class63.go` - ORM Implementation

### Advanced Topics
- `class70.go` - Concurrency with Goroutines
- `class71.go` - Channels and Select
- `class72.go` - Context and Cancellation
- `class73.go` - Testing and Benchmarks

### Microservices
- `class80.go` - gRPC Server
- `class81.go` - Message Queues
- `class82.go` - Docker Integration
- `class83.go` - Health Checks

## 📁 Complete Directory Structure

```
Backend/
├── class47.go                    # Basic HTTP Server
├── class48.go                    # User Management API
├── class49.go                    # Your next class
├── class50.go                    # Another class
├── generate-class-doc.sh         # Documentation generator
├── go.mod                        # Go module file
├── README.md                     # Project README
├── WORKFLOW_README.md            # This detailed guide
├── SETUP_COMMANDS.md             # Git setup instructions
├── .github/
│   └── workflows/
│       └── docs.yml              # GitHub Actions for auto-deployment
└── docs/                         # Hugo documentation site
    ├── hugo.toml                 # Hugo configuration
    ├── content/
    │   ├── docs/
    │   │   ├── _index.md         # Main documentation page
    │   │   ├── classes/          # Auto-generated class documentation
    │   │   │   ├── _index.md     # Classes listing page
    │   │   │   ├── class47.md    # Class 47 documentation
    │   │   │   ├── class48.md    # Class 48 documentation
    │   │   │   └── class49.md    # Auto-generated when you run script
    │   │   ├── api/
    │   │   │   └── reference.md  # API reference
    │   │   └── project-overview/
    │   │       ├── _index.md
    │   │       └── intro.md
    │   └── menu/
    │       └── index.md          # Navigation menu
    ├── public/                   # Generated website (auto-created)
    └── themes/
        └── book/                 # Hugo Book theme
```

## 🔧 Advanced Tips

### Adding Rich Comments

```go
// Package main demonstrates advanced HTTP server concepts
package main

// UserService provides user management functionality
// It handles CRUD operations and authentication for users
type UserService struct {
    users    []User         // In-memory user storage
    tokens   map[string]int // Active session tokens
    lastID   int            // Last assigned user ID
}

// CreateUser adds a new user to the system
// Parameters:
//   - name: User's full name
//   - email: User's email address (must be unique)
// Returns:
//   - User: Created user object
//   - error: Error if creation fails
func (us *UserService) CreateUser(name, email string) (User, error) {
    // Implementation here
}
```

### Environment-Specific Configuration

```go
// Config holds application configuration
type Config struct {
    Port     string `json:"port"`
    Debug    bool   `json:"debug"`
    Database string `json:"database"`
}

// LoadConfig loads configuration from environment or defaults
func LoadConfig() Config {
    return Config{
        Port:     getEnv("PORT", "8080"),
        Debug:    getEnv("DEBUG", "false") == "true",
        Database: getEnv("DB_URL", "localhost:5432"),
    }
}
```

## 🚀 Quick Start Checklist

- [ ] Choose your class number (next in sequence)
- [ ] Create `classXX.go` file with proper structure
- [ ] Add descriptive comments
- [ ] Use unique function names (avoid `main`, use `ClassXXMain`)
- [ ] Test your code locally
- [ ] Run `./generate-class-doc.sh classXX.go`
- [ ] Commit and push: `git add . && git commit -m "Add classXX.go: description" && git push origin main`
- [ ] Check your documentation website updates automatically!

## 🌐 Your Documentation Website

After pushing, your documentation will be available at:
`https://mdimamhosen.github.io/ecom-go/`

The website includes:
- ✅ **Homepage** with project overview
- ✅ **Classes section** with all your Go files
- ✅ **API reference** with endpoint documentation
- ✅ **Search functionality** to find content quickly
- ✅ **Mobile-responsive design**
- ✅ **Automatic updates** when you push new files

---

**Happy coding!** 🎉 Every new Go file you create will automatically get beautiful documentation on your website!
