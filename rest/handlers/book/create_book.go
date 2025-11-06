package book

import (
	"bookShop/repo/book"
	"bookShop/util"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // 10 MB max
	title := r.FormValue("title")
	author := r.FormValue("author")
	priceStr := r.FormValue("price")
	price, _ := strconv.ParseFloat(priceStr, 32)
	description := r.FormValue("description")
	category := r.FormValue("category")
	is_stock := r.FormValue("is_stock")
	stock, _ := strconv.ParseBool(is_stock)
	file, handler, err := r.FormFile("image_path")

	var filePath string
	if err == nil {
		defer file.Close()
		os.MkdirAll("uploads", os.ModePerm)
		filePath = filepath.Join("uploads", handler.Filename)
		filePath = filepath.ToSlash(filePath)
		dst, _ := os.Create(filePath)
		defer dst.Close()
		io.Copy(dst, file)
	}

	// Create new book
	createdBook, err := h.bookRepo.Create(book.Book{
		Title:       title,
		Author:      author,
		Price:       float32(price),
		Description: description,
		ImagePath:   filePath,
		Category:    category,
		IsStock:     stock,
	})

	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, createdBook, http.StatusCreated)
}
