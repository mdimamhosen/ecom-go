package cmd

import (
	"ecommerce/global"
	"ecommerce/handlers"
	"fmt"
	"log"
	"net/http"
)

func Server() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /createProduct", http.HandlerFunc(handlers.CreateProduct))

	// CORS middleware handles all OPTIONS globally
	handler := global.CorsMiddleware(mux)

	fmt.Println("Server is running on port 8080")
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
