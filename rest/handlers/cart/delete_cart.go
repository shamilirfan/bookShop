package cart

import (
	"bookShop/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteCart(w http.ResponseWriter, r *http.Request) {
	str := r.PathValue("id")
	id, err := strconv.Atoi(str)

	if err != nil {
		util.SendError(w, "Invalid id", http.StatusBadRequest)
		return
	}

	err = h.cartRepo.Delete(id)
	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, "Sccessfully Deleted", http.StatusOK)
}
