package book

import (
	"bookShop/repo/book"
	"bookShop/util"
	"encoding/json"
	"net/http"
	"strconv"
)

type UpBook struct {
	ID           int     `json:"id" db:"id"` // It is called tag
	Title        string  `json:"title" db:"title"`
	Author       string  `json:"author" db:"author"`
	Price        float32 `json:"price" db:"price"`
	Description  string  `json:"description" db:"description"`
	ImageUrl     string  `json:"image_url" db:"image_url"`
	BookCategory string  `json:"book_catagory" db:"book_catagory"`
	IsStock      bool    `json:"is_stock" db:"is_stock"`
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
	h.bookRepo.Update(book.Book{
		ID:           id,
		Title:        updatedBook.Title,
		Author:       updatedBook.Author,
		Price:        updatedBook.Price,
		Description:  updatedBook.Description,
		ImageUrl:     updatedBook.ImageUrl,
		BookCategory: updatedBook.BookCategory,
		IsStock:      updatedBook.IsStock,
	}) // Call update function

	util.SendData(w, "Successfully updated", http.StatusOK) // Call sendData function
}
