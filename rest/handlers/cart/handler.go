package cart

import (
	"bookShop/repo/cart"
)

type Handler struct {
	cartRepo cart.CartRepo
}

func NewHandler(cartRepo cart.CartRepo) *Handler {
	return &Handler{cartRepo: cartRepo}
}
