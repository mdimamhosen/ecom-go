package rest

import (
	handlers "ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/users"
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func iniRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// Product APIs
	mux.Handle("GET /products", manager.WrapMux(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductByID))
	mux.Handle("PUT /products/{id}", http.HandlerFunc(handlers.UpdateProduct))
	//TODO: Complete delete route and wrap with manager

	// User APIs
	mux.Handle("POST /users", manager.WrapMux(http.HandlerFunc(users.CreateUser)))
}
