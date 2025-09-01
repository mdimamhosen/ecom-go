package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Welcome to the home page! This is a simple HTTP server in Go.")
}

type TProduct struct {
	ID          int
	TITLE       string
	DESCRIPTION string
	PRICE       float64
	ImgUrl      string
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Controll-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		http.Error(w, "Please give me get request", 400)
		return
	}

	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.Encode(productList)
}

var productList []TProduct

func main() {

	mux := http.NewServeMux() // router

	mux.HandleFunc("/", homeHandler) // home route

	mux.HandleFunc("/products", getProducts)

	fmt.Println("Server is running on port 3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}

}

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
