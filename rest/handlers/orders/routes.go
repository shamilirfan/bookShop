package orders

import (
	"bookShop/rest/middlewares"
	"net/http"
)

func (h *Handler) RregisterRoutes(mux *http.ServeMux) {
	// Create route/endpoint
	mux.Handle("POST /orders",
		middlewares.Use(http.HandlerFunc(h.CreateOrder)))
	mux.Handle("PUT /orders/{id}",
		middlewares.Use(http.HandlerFunc(h.UpdateOrder)))
}
