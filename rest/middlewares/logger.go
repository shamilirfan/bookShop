package middlewares

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r) // এটা সবসময় “পরের handler কে কল করো”।
		end := time.Since(start)

		log.Print(r.Method, r.URL, end)
	})
}
