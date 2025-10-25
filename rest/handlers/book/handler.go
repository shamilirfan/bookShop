package book

import "bookShop/repo"

type Handler struct {
	bookRepo repo.BookRepo
}

func NewHandler(bookRepo repo.BookRepo) *Handler {
	return &Handler{
		bookRepo: bookRepo,
	}
}
