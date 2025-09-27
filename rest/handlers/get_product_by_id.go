package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
	"strconv"
)

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, "Please Provide a valid id")
		return
	}
	product := database.Get(id)

	if product.ID == 0 {
		utils.SendError(w, 404, "Product not found")
		return
	}

	utils.SendJSON(w, 201, product)

}
