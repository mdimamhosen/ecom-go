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
	productList = []Product{}
	StoreMu     sync.RWMutex
)

func Store(p Product) []Product {
	p.ID = len(productList) + 1
	productList = append(productList, p)
	return productList
}

func List() []Product {
	return productList
}

func Get(id int) Product {
	for _, product := range productList {
		if id == product.ID {
			return product
		}
	}
	return Product{}
}

func Update(product Product) Product {
	for idx, p := range productList {
		if p.ID == product.ID {
			productList[idx] = product

			return product
		}
	}
	return Product{}
}

func Delete(productId int) {
	var tempList []Product

	for _, p := range productList {
		if p.ID != productId {
			tempList = append(tempList, p)
		}
	}
	productList = tempList
}

func init() {
	prod1 := Product{
		ID:          1,
		Title:       "Orange",
		Description: "Very Sweet fruit",
		Price:       100,
		ImgURL:      "https://images.unsplash.com/photo-1557800636-894a64c1696f?w=500&auto=format&fit=crop&q=60&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxzZWFyY2h8Mnx8b3JhbmdlfGVufDB8fDB8fHww",
	}

	productList = append(productList, prod1)
}
