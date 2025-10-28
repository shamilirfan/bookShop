package users

import (
	"bookShop/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.Handle("POST /signup",
		middlewares.Use(http.HandlerFunc(h.SignUp)),
	)
	mux.Handle("POST /signin",
		middlewares.Use(http.HandlerFunc(h.SignIn)),
	)
}
