package book

import (
	"bookShop/util"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Path value
	bookID := r.PathValue("id")
	id, err := strconv.Atoi(bookID)

	// Error handling
	if err != nil {
		util.SendError(w, "Please give a valid id", http.StatusBadRequest)
		return
	}

	h.bookRepo.Delete(id)                                   // Call delete function
	util.SendData(w, "Successfully deleted", http.StatusOK) // Call sendData function
}
