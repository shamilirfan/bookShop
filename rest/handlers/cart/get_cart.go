package cart

import (
	"bookShop/util"
	"net/http"
)

func (h *Handler) GetCart(w http.ResponseWriter, r *http.Request) {
	cartList, err := h.cartRepo.Get()
	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, cartList, http.StatusOK)
}
