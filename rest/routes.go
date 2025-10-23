package rest

import (
	"bookShop/rest/handlers/book"
	"bookShop/rest/middlewares"
	"net/http"
)

func Routes(mux *http.ServeMux) {
	// Create route/endpoint
	mux.Handle("GET /books",
		middlewares.Use(http.HandlerFunc(book.GetBooks)))
	mux.Handle("GET /books/{id}",
		middlewares.Use(http.HandlerFunc(book.GetBook)))
	mux.Handle("POST /books",
		middlewares.Use(http.HandlerFunc(book.CreateBook)))
	mux.Handle("PUT /books/{id}",
		middlewares.Use(http.HandlerFunc(book.UpdateBook)))
	mux.Handle("DELETE /books/{id}",
		middlewares.Use(http.HandlerFunc(book.DeleteBook)))
}
