package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Please send a POST request")
		return
	}

	var newProd database.Product
	if err := json.NewDecoder(r.Body).Decode(&newProd); err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	// simple validation
	if newProd.Title == "" {
		utils.WriteError(w, http.StatusBadRequest, "title is required")
		return
	}

	database.StoreMu.Lock()
	defer database.StoreMu.Unlock()

	newProd.ID = len(database.ProductList) + 1
	database.ProductList = append(database.ProductList, newProd)

	utils.SendJSON(w, http.StatusCreated, newProd)
}
