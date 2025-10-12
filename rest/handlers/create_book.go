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

	// Write header
	w.WriteHeader(http.StatusCreated)

	// Write a new book's ID
	newBook.ID = len(database.BookList) + 1

	// Append new book in a book list
	database.BookList = append(database.BookList, newBook)

	// Encode
	json.NewEncoder(w).Encode(&newBook)
}
