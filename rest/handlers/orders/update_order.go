package orders

import (
	"bookShop/repo/orders"
	"bookShop/util"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	str := r.PathValue("id")
	id, _ := strconv.Atoi(str)

	var order orders.Order
	err := json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		util.SendError(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	_, err = h.orderRepo.Update(orders.Order{
		ID:     id,
		Status: order.Status,
	})
	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, "✅Successfully Status Updated", http.StatusOK)
}
