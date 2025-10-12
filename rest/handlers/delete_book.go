package handlers

import (
	"bookShop/database"
	"bookShop/util"
	"net/http"
	"strconv"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Path value
	bookID := r.PathValue("id")
	id, err := strconv.Atoi(bookID)

	// Error handling
	if err != nil {
		http.Error(w, "Please give a valid id", http.StatusBadRequest)
		return
	}

	// Store unmatched value
	var tempList []database.Book

	// Searching specific value
	for _, value := range database.BookList {
		if id != value.ID {
			tempList = append(tempList, value)
		}
	}

	database.BookList = tempList

	util.SendData(w, "Successfully deleted", http.StatusOK)
}
