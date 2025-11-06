package orders

import (
	"bookShop/util"
	"net/http"
)

func (h *Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	orderList, err := h.orderRepo.Get()

	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, orderList, http.StatusOK)
}
