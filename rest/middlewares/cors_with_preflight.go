package middlewares

import (
	"bookShop/config"
	"net/http"
	"strconv"
)

func CorsWithPreflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/* Cors handling
		Allowed origin = frontend origin
		*/
		str := strconv.Itoa(int(config.GetConfig().HttpPort))
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:"+str)
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		// Allowed headers
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Allowed methods
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

		// Handle preflight OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Call next middleware/handler
		next.ServeHTTP(w, r)
	})
}
