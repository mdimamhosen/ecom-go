package handlers

import (
	"ecommerce/database"
	"ecommerce/utils"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		utils.WriteError(w, http.StatusMethodNotAllowed, "Please send a GET request")
		return
	}

	// database.StoreMu.RLock()
	// defer database.StoreMu.RUnlock()
	product := database.List()
	utils.SendJSON(w, http.StatusOK, product)
}
