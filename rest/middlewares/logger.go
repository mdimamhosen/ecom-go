package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// recover from panics in downstream handlers so the server doesn't crash
		// defer func() {
		// 	if rec := recover(); rec != nil {
		// 		log.Printf("panic recovered in middleware.Logger: %v\n", rec)
		// 		// best-effort: write a 500 response if possible
		// 		http.Error(w, "internal server error", http.StatusInternalServerError)
		// 	}
		// 	log.Println("After next call...")
		// 	log.Println(r.Method, r.URL.Path, time.Since(start))
		// 	log.Println("........")
		// }()

		next.ServeHTTP(w, r)
		log.Println(r.Method, r.URL.Path, time.Since(start))

	})
}
