package handlers

import (
	"bookShop/database"
	"bookShop/util"
	"encoding/json"
	"net/http"
	"strconv"
)

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Path value
	bookID := r.PathValue("id")
	id, err := strconv.Atoi(bookID)

	// Error handling
	if err != nil {
		http.Error(w, "Please give me a valid id", http.StatusBadRequest)
		return
	}

	// Store updated book
	var updatedBook database.Book

	// Decode
	decoder := json.NewDecoder(r.Body).Decode(&updatedBook)

	// Error handling
	if decoder != nil {
		http.Error(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	// Searching specific book
	for index, value := range database.BookList {
		if id == value.ID {
			updatedBook.ID = id
			database.BookList[index] = updatedBook
		}
	}

	util.SendData(w, "Successfully updated", http.StatusOK)
}
