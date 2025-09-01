package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// User represents a user in our system
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Active  bool   `json:"active"`
	Created string `json:"created"`
}

// UserService handles user-related operations
type UserService struct {
	users []User
}

// NewUserService creates a new user service instance
func NewUserService() *UserService {
	return &UserService{
		users: []User{
			{ID: 1, Name: "John Doe", Email: "john@example.com", Active: true, Created: "2025-09-01"},
			{ID: 2, Name: "Jane Smith", Email: "jane@example.com", Active: true, Created: "2025-09-01"},
		},
	}
}

// GetAllUsers returns all users
func (us *UserService) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(us.users)
}

// GetUserByID returns a specific user by ID
func (us *UserService) GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Extract ID from URL path
	path := strings.TrimPrefix(r.URL.Path, "/users/")
	id, err := strconv.Atoi(path)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Find user
	for _, user := range us.users {
		if user.ID == id {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

// class48HomeHandler handles the root endpoint for class48
func class48HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{
		"message":   "Welcome to Class 48 - User Management API",
		"version":   "1.0.0",
		"endpoints": "/users, /users/{id}",
	}

	json.NewEncoder(w).Encode(response)
}

// Class48Main demonstrates the user management server
// This is a separate main function for demonstration
func Class48Main() {
	// Initialize user service
	userService := NewUserService()

	// Create router
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/", class48HomeHandler)
	mux.HandleFunc("/users", userService.GetAllUsers)
	mux.HandleFunc("/users/", userService.GetUserByID)

	// Server configuration
	port := ":8080"
	fmt.Printf("🚀 Class 48 Server starting on port %s\n", port)
	fmt.Println("📖 Available endpoints:")
	fmt.Println("   GET /          - API information")
	fmt.Println("   GET /users     - Get all users")
	fmt.Println("   GET /users/1   - Get user by ID")

	// Start server
	log.Fatal(http.ListenAndServe(port, mux))
}
