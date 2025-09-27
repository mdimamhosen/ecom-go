package rest

import (
	"ecommerce/rest/handlers"
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func iniRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("GET /products", manager.WrapMux(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductByID))

	mux.Handle("PUT /products/{id}", http.HandlerFunc(handlers.UpdateProduct))
}
