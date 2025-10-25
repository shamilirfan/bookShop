package book

import (
	"bookShop/util"
	"net/http"
)

func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	bookList, err := h.bookRepo.List()

	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
	}

	util.SendData(w, bookList, http.StatusOK)
}
