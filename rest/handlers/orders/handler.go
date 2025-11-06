package orders

import "bookShop/repo/orders"

type Handler struct {
	orderRepo orders.OrderRepo
}

func NewHandler(orderRepo orders.OrderRepo) *Handler {
	return &Handler{orderRepo: orderRepo}
}
