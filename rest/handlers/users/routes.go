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
	mux.Handle("POST /request-reset",
		middlewares.Use(http.HandlerFunc(h.Request_Password_Reset)),
	)

	mux.Handle("GET /reset-password-form",
		middlewares.Use(http.HandlerFunc(h.Reset_Password_Form)),
	)

	mux.Handle("POST /reset-password",
		middlewares.Use(http.HandlerFunc(h.Reset_Password)),
	)
}
