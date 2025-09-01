package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// ClassXX demonstrates [DESCRIBE WHAT THIS CLASS TEACHES]
// Replace XX with your class number and update the description

// [YourStruct] represents [what it represents in your system]
type YourStruct struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	// Add more fields as needed
}

// [YourService] handles [what operations it handles]
type YourService struct {
	data []YourStruct // Your data storage
}

// NewYourService creates a new instance of YourService
func NewYourService() *YourService {
	return &YourService{
		data: []YourStruct{
			{ID: 1, Name: "Sample Item 1"},
			{ID: 2, Name: "Sample Item 2"},
		},
	}
}

// GetAll returns all items (HTTP handler example)
func (ys *YourService) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(ys.data)
}

// classXXHomeHandler handles the root endpoint for this class
func classXXHomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := map[string]string{
		"message":     "Welcome to Class XX - [Your Topic]",
		"description": "[Brief description of what this class demonstrates]",
		"endpoints":   "[List your available endpoints]",
	}

	json.NewEncoder(w).Encode(response)
}

// ClassXXMain demonstrates the main functionality of this class
// Replace XX with your actual class number
func ClassXXMain() {
	service := NewYourService()

	mux := http.NewServeMux()
	mux.HandleFunc("/", classXXHomeHandler)
	mux.HandleFunc("/items", service.GetAll)

	port := ":8080"
	fmt.Printf("🚀 Class XX Server starting on port %s\n", port)
	fmt.Println("📖 Available endpoints:")
	fmt.Println("   GET /       - Class information")
	fmt.Println("   GET /items  - Get all items")
	fmt.Println("📝 This class demonstrates: [Your learning objective]")

	// In a real scenario, you'd run this:
	// log.Fatal(http.ListenAndServe(port, mux))

	// For template purposes, just show the setup:
	fmt.Println("✅ Server setup complete (template mode)")
}

// ClassXXDemo runs a demonstration of this class
func ClassXXDemo() {
	fmt.Println("🎯 Class XX Demo: [Your Topic]")
	fmt.Println("This function demonstrates the key concepts...")

	// Add your demo code here
	service := NewYourService()
	fmt.Printf("📊 Created service with %d items\n", len(service.data))

	ClassXXMain()
}

/*
USAGE INSTRUCTIONS:
1. Copy this template to a new file: classXX.go (replace XX with your number)
2. Replace all instances of:
   - XX with your class number
   - YourStruct with your actual struct name
   - YourService with your actual service name
   - [PLACEHOLDER TEXT] with your actual descriptions
3. Implement your specific logic
4. Run: ./generate-class-doc.sh classXX.go
5. Commit and push to GitHub

EXAMPLES OF WHAT TO REPLACE:
- "ClassXX demonstrates" → "Class49 demonstrates JWT Authentication"
- "YourStruct" → "AuthToken" or "User" or "Product"
- "YourService" → "AuthService" or "UserService" or "ProductService"
- "[Your Topic]" → "JWT Authentication" or "Database Operations"
*/
