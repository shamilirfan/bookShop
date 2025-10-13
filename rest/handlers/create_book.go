package handlers

import (
	"bookShop/database"
	"encoding/json"
	"net/http"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	// Store new book
	var newBook database.Book

	// Decode
	decoder := json.NewDecoder(r.Body).Decode(&newBook)

	// Error handling
	if decoder != nil {
		http.Error(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)   // Write header
	database.Create(newBook)            // Call Create function
	json.NewEncoder(w).Encode(&newBook) // Encode
}
