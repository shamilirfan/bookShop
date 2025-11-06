package orders

import (
	"bookShop/repo/orders"
	"bookShop/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var ord orders.Order

	err := json.NewDecoder(r.Body).Decode(&ord)
	if err != nil {
		util.SendError(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	order, err := h.orderRepo.Create(orders.Order{
		UserID:        ord.UserID,
		RoadNumber:    ord.RoadNumber,
		HoldingNumber: ord.HoldingNumber,
		Area:          ord.Area,
		Thana:         ord.Thana,
		District:      ord.District,
		PhoneNumber:   ord.PhoneNumber,
		Status:        ord.Status,
		Items:         ord.Items,
	})

	if err != nil {
		util.SendError(w, err, http.StatusInternalServerError)
		return
	}

	util.SendData(w, order, http.StatusCreated)
}
