package book

import (
	"bookShop/repo"
	"bookShop/util"
	"encoding/json"
	"net/http"
	"strconv"
)

type UpBook struct {
	ID           int     `json:"id"` // It is called tag
	Title        string  `json:"title"`
	Author       string  `json:"author"`
	Price        float32 `json:"price"`
	Description  string  `json:"description"`
	ImageUrl     string  `json:"imageUrl"`
	BookCatagory string  `json:"bookCatagory"`
	IsStock      bool    `json:"isStock"`
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	// Path value
	bookID := r.PathValue("id")
	id, err := strconv.Atoi(bookID)

	// Error handling
	if err != nil {
		util.SendError(w, "Please give me a valid id", http.StatusBadRequest)
		return
	}

	// Store updated book
	var updatedBook UpBook

	// Decode
	decoder := json.NewDecoder(r.Body).Decode(&updatedBook)

	// Error handling
	if decoder != nil {
		util.SendError(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	updatedBook.ID = id // Update id
	h.bookRepo.Update(repo.Book{
		ID:           id,
		Title:        updatedBook.Title,
		Author:       updatedBook.Author,
		Price:        updatedBook.Price,
		Description:  updatedBook.Description,
		ImageUrl:     updatedBook.ImageUrl,
		BookCatagory: updatedBook.BookCatagory,
		IsStock:      updatedBook.IsStock,
	}) // Call update function

	util.SendData(w, "Successfully updated", http.StatusOK) // Call sendData function
}
