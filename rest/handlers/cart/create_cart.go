package cart

import (
	"bookShop/repo/cart"
	"bookShop/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateCart(w http.ResponseWriter, r *http.Request) {
	var cart cart.Cart

	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		util.SendError(w, "Invalid input", http.StatusBadRequest)
		return
	}

	newCart, err := h.cartRepo.Create(cart)
	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, newCart, http.StatusCreated)
}
