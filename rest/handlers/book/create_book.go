package book

import (
	"bookShop/database"
	"bookShop/util"
	"encoding/json"
	"net/http"
)

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	// Store new book
	var newBook database.Book

	// Decode
	decoder := json.NewDecoder(r.Body).Decode(&newBook)

	// Error handling
	if decoder != nil {
		util.SendError(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	createdBook := database.Create(newBook)           // Call Create function
	util.SendData(w, createdBook, http.StatusCreated) // Encode
}
