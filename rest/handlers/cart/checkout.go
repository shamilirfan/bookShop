package cart

import (
	"bookShop/repo/cart"
	"bookShop/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) Checkout(w http.ResponseWriter, r *http.Request) {
	var checkout cart.CheckoutRequest
	json.NewDecoder(r.Body).Decode(&checkout)

	err := h.cartRepo.Checkout(checkout)
	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, "Order placed", http.StatusOK)
}
