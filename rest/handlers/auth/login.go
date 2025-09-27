package auth

import (
	"ecommerce/database"
	"ecommerce/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LogIn(w http.ResponseWriter, r *http.Response) {
	var reqLogin ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqLogin)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request data", http.StatusBadRequest)
	}

	user := database.Find(reqLogin.Email, reqLogin.Password)
	if user == nil {
		http.Error(w, "Invalid credentials", http.StatusBadRequest)
	}

	utils.SendJSON(w, 201, user)

}
