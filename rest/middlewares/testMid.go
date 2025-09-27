package middleware

import (
	"log"
	"net/http"
)

func TestMid(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Test middleware...")

		next.ServeHTTP(w, r)
	})
}
