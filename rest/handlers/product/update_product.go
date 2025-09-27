package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	pid := r.PathValue("id")

	id, err := strconv.Atoi(pid)

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Please Provide a valid id")
		return
	}

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Provide a valid json")
		return
	}

	newProduct.ID = id

	product := database.Update(newProduct)

	utils.SendJSON(w, 201, product)
}
