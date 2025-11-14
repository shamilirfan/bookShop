package book

import (
	"bookShop/util"
	"net/http"
)

func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	bookList, err := h.bookRepo.Get()
	if err != nil {
		util.SendError(w, "Failed to fetch books", http.StatusInternalServerError)
		return
	}

	util.SendData(w, bookList, http.StatusOK)
}
