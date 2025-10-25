package book

import (
	"bookShop/util"
	"net/http"
	"strconv"
)

func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	// Path Value
	bookID := r.PathValue("id")
	id, err := strconv.Atoi(bookID)

	// Error handling
	if err != nil {
		util.SendError(w, "Please give me a valid id", http.StatusBadRequest)
		return
	}

	book, err := h.bookRepo.GetByID(id) // Call GetByID function

	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if book != nil {
		util.SendData(w, book, http.StatusOK)
		return
	}

	// Book not found
	util.SendData(w, "Book Not Found!!!", http.StatusNotFound)
}
