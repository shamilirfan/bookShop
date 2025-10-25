package book

import (
	"bookShop/repo"
	"bookShop/util"
	"encoding/json"
	"net/http"
)

// Struct define
type CreateBook struct {
	ID           int     `json:"id"` // It is called tag
	Title        string  `json:"title"`
	Author       string  `json:"author"`
	Price        float32 `json:"price"`
	Description  string  `json:"description"`
	ImageUrl     string  `json:"imageUrl"`
	BookCatagory string  `json:"bookCatagory"`
	IsStock      bool    `json:"isStock"`
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	// Store new book
	var newBook CreateBook

	// Decode
	decoder := json.NewDecoder(r.Body).Decode(&newBook)

	// Error handling
	if decoder != nil {
		util.SendError(w, "Please give me a valid json", http.StatusBadRequest)
		return
	}

	createdBook, err := h.bookRepo.Create(repo.Book{
		Title:        newBook.Title,
		Author:       newBook.Author,
		Price:        newBook.Price,
		Description:  newBook.Description,
		ImageUrl:     newBook.ImageUrl,
		BookCatagory: newBook.BookCatagory,
		IsStock:      newBook.IsStock,
	}) // Call Create function

	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
	}

	util.SendData(w, createdBook, http.StatusCreated) // Encode
}
