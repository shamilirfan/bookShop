package book

import (
	"bookShop/util"
	"log"
	"net/http"
)

func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	bookList, err := h.bookRepo.Get()
	if err != nil {
		log.Printf("Failed to fetch books: %v", err)
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	log.Printf("Fetched %d books", len(bookList))
	util.SendData(w, bookList, http.StatusOK)
}
