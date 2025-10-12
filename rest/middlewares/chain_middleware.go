package middlewares

import (
	"net/http"
)

type Middleware func(http.Handler) http.Handler

var middlewareList []Middleware

func ChainMiddleware(middlewares ...Middleware) {
	middlewareList = append(middlewareList, middlewares...)
}

func Use(handler http.Handler) http.Handler {
	for i := 0; i < len(middlewareList); i++ {
		handler = middlewareList[i](handler)
	}
	return handler
}
