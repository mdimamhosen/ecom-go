package utils

import (
	"encoding/json"
	"net/http"
)

func SendJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func SendError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
}
