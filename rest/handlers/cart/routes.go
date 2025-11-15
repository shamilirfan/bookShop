package cart

import (
	"bookShop/rest/middlewares"
	"net/http"
)

func (h *Handler) RregisterRoutes(mux *http.ServeMux) {
	// Create route/endpoint
	mux.Handle("GET /cart",
		middlewares.Use(http.HandlerFunc(h.GetCart)))
	mux.Handle("POST /cart",
		middlewares.Use(http.HandlerFunc(h.CreateCart)))
	mux.Handle("POST /checkout",
		middlewares.Use(http.HandlerFunc(h.Checkout)))
	mux.Handle("DELETE /checkout/{id}",
		middlewares.Use(http.HandlerFunc(h.DeleteCart)))
}
