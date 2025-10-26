package users

import (
	"bookShop/util"
	"net/http"
)

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, users, http.StatusOK)
}
