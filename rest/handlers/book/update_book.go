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
		util.SendError(w, "Please give me a valid id", http.StatusBadRequest)
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

	updatedBook.ID = id                                     // Update id
	database.Update(updatedBook)                            // Call update function
	util.SendData(w, "Successfully updated", http.StatusOK) // Call sendData function
}
