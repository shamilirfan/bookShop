package book

import (
	"bookShop/repo"
	"bookShop/util"
	"encoding/json"
	"net/http"
)

// Struct define
type ReqCreateBook struct {
	ID           int     `json:"id" db:"id"`
	Title        string  `json:"title" db:"title"`
	Author       string  `json:"author" db:"author"`
	Price        float32 `json:"price" db:"price"`
	Description  string  `json:"description" db:"description"`
	ImageUrl     string  `json:"imageUrl" db:"image_url"`
	BookCategory string  `json:"bookCategory" db:"book_category"`
	IsStock      bool    `json:"isStock" db:"is_stock"`
}

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook ReqCreateBook

	// Decode JSON request
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		util.SendError(w, "Please give me a valid JSON", http.StatusBadRequest)
		return
	}

	// Create new book
	createdBook, err := h.bookRepo.Create(repo.Book{
		Title:        newBook.Title,
		Author:       newBook.Author,
		Price:        newBook.Price,
		Description:  newBook.Description,
		ImageUrl:     newBook.ImageUrl,
		BookCategory: newBook.BookCategory,
		IsStock:      newBook.IsStock,
	})

	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, createdBook, http.StatusCreated)
}
