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

	// Store specific book
	var book database.Book

	// searching book
	for _, value := range database.BookList {
		if id == value.ID {
			book = value

			util.SendData(w, &book, http.StatusOK)
			return
		}
	}

	// Book not found
	util.SendData(w, "Book Not Found!!!", http.StatusNotFound)
}
