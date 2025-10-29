package book

import (
	"bookShop/repo/book"
	"bookShop/util"
	"encoding/json"
	"net/http"
	"strconv"
)

type UpBook struct {
	ID          int     `json:"id" db:"id"` // It is called tag
	Title       string  `json:"title" db:"title"`
	Author      string  `json:"author" db:"author"`
	Price       float32 `json:"price" db:"price"`
	Description string  `json:"description" db:"description"`
	ImagePath   string  `json:"image_path" db:"image_path"`
	Category    string  `json:"catagory" db:"catagory"`
	IsStock     bool    `json:"is_stock" db:"is_stock"`
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
		ID:          id,
		Title:       updatedBook.Title,
		Author:      updatedBook.Author,
		Price:       updatedBook.Price,
		Description: updatedBook.Description,
		ImagePath:   updatedBook.ImagePath,
		Category:    updatedBook.Category,
		IsStock:     updatedBook.IsStock,
	}) // Call update function

	util.SendData(w, "Successfully updated", http.StatusOK) // Call sendData function
}
