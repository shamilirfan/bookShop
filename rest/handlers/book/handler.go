package book

import "bookShop/repo/book"

type Handler struct {
	bookRepo book.BookRepo
}

func NewHandler(bookRepo book.BookRepo) *Handler {
	return &Handler{bookRepo: bookRepo}
}
