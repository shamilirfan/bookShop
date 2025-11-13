package book

import (
	"bookShop/repo/book"
	"github.com/cloudinary/cloudinary-go/v2"
)

type Handler struct {
	bookRepo book.BookRepo
	cld      *cloudinary.Cloudinary
}

func NewHandler(bookRepo book.BookRepo, cld *cloudinary.Cloudinary) *Handler {
	return &Handler{bookRepo: bookRepo, cld: cld}
}
