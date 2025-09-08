package database

import "sync"

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImgURL      string  `json:"imgUrl"`
}

// in-memory store + mutex for concurrency safety
var (
	ProductList = []Product{}
	StoreMu     sync.RWMutex
)

func init() {
	prod1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Very Sweet fruit",
		Price:       100,
		ImgURL:      "https://images.unsplash.com/photo-1557800636-894a64c1696f?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8b3JhbmdlfGVufDB8fDB8fHww",
	}

	ProductList = append(ProductList, prod1)
}
