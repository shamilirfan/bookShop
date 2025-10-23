package book

import (
	"bookShop/rest/middlewares"
	"net/http"
)

func (h *Handler) RregisterRoutes(mux *http.ServeMux) {
	// Create route/endpoint
	mux.Handle("GET /books",
		middlewares.Use(http.HandlerFunc(h.GetBooks)))
	mux.Handle("GET /books/{id}",
		middlewares.Use(http.HandlerFunc(h.GetBook)))
	mux.Handle("POST /books",
		middlewares.Use(http.HandlerFunc(h.CreateBook)))
	mux.Handle("PUT /books/{id}",
		middlewares.Use(http.HandlerFunc(h.UpdateBook)))
	mux.Handle("DELETE /books/{id}",
		middlewares.Use(http.HandlerFunc(h.DeleteBook)))
}
