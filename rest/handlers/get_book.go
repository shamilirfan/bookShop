package handlers

import (
	"bookShop/database"
	"bookShop/util"
	"net/http"
	"strconv"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	// Path Value
	bookID := r.PathValue("id")
	id, err := strconv.Atoi(bookID)

	// Error handling
	if err != nil {
		http.Error(w, "Please give me a valid id", http.StatusBadRequest)
		return
	}

	book := database.GetByID(id) // Call GetByID function

	if book != nil {
		util.SendData(w, book, http.StatusOK)
	}

	// Book not found
	util.SendData(w, "Book Not Found!!!", http.StatusNotFound)
}
