---
title: "class47"
weight: 100
date: 2025-09-01T23:46:20+0600
---

# class47 Documentation

## Source Code

```go
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
```

## File Information

- **Filename**: class47.go
- **Created**: Mon 01 Sep 2025 11:46:20 PM +06
- **Size**: 1532 bytes
- **Lines**: 65 lines

## Functions and Types

- Line 9:func homeHandler(w http.ResponseWriter, r *http.Request) {
- Line 15:type TProduct struct {
- Line 23:func getProducts(w http.ResponseWriter, r *http.Request) {
- Line 38:func main() {
- Line 55:func init() {

## Description



---

*This documentation was auto-generated from the source file.*
