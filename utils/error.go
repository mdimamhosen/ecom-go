package utils

import "net/http"

func WriteError(w http.ResponseWriter, status int, msg string) {
	SendJSON(w, status, map[string]string{"error": msg})
}
