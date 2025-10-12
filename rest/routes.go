package rest

import (
	"bookShop/rest/handlers"
	"bookShop/rest/middlewares"
	"net/http"
)

func Routes(mux *http.ServeMux) {
	// Create route/endpoint
	mux.Handle("GET /books",
		middlewares.Use(http.HandlerFunc(handlers.GetBooks)))
	mux.Handle("GET /books/{id}",
		middlewares.Use(http.HandlerFunc(handlers.GetBook)))
	mux.Handle("POST /books",
		middlewares.Use(http.HandlerFunc(handlers.CreateBook)))
	mux.Handle("PUT /books/{id}",
		middlewares.Use(http.HandlerFunc(handlers.UpdateBook)))
	mux.Handle("DELETE /books/{id}",
		middlewares.Use(http.HandlerFunc(handlers.DeleteBook)))
}
