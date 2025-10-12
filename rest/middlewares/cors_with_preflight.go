package middlewares

import (
	"net/http"
)

func CorsWithPreflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Cors handling
		w.Header().Set("Access-Control-Allow-Origin", "*") // * means any IP address can access
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
