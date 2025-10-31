package book

import (
	"bookShop/repo/book"
	"bookShop/util"
	"database/sql"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	err1 := r.ParseMultipartForm(10 << 20)
	bookID := r.PathValue("id")
	id, err2 := strconv.Atoi(bookID)
	title := r.FormValue("title")
	author := r.FormValue("author")
	priceStr := r.FormValue("price")
	price, err3 := strconv.ParseFloat(priceStr, 64)
	description := r.FormValue("description")
	file, handler, err4 := r.FormFile("image_path")
	category := r.FormValue("category")
	isStockStr := r.FormValue("is_stock")
	is_stock, err5 := strconv.ParseBool(isStockStr)

	// ParseMultipartForm error handling
	if err1 != nil {
		util.SendError(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// book ID error handling
	if err2 != nil {
		util.SendError(w, "Please give a valid book ID", http.StatusBadRequest)
		return
	}

	// price error handling
	if err3 != nil {
		util.SendError(w, "Price must be a number", http.StatusBadRequest)
		return
	}

	// is_stock  error handling
	if err5 != nil {
		util.SendError(w, "isStock must be a boolean", http.StatusBadRequest)
		return
	}

	// image_path  error handling
	var filePath string
	if err4 == nil {
		defer file.Close()

		os.MkdirAll("uploads", os.ModePerm)
		filePath = filepath.Join("uploads", handler.Filename)
		filePath = filepath.ToSlash(filePath)

		dst, err6 := os.Create(filePath)
		if err6 != nil {
			util.SendError(w, "Failed to save image", http.StatusInternalServerError)
			return
		}
		defer dst.Close()

		_, err7 := io.Copy(dst, file)
		if err7 != nil {
			util.SendError(w, "Failed to copy image content", http.StatusInternalServerError)
			return
		}

	} else {
		util.SendError(w, "Could not read the file", http.StatusBadRequest)
		return
	}

	// Update book
	updatedBook, err8 := h.bookRepo.Update(book.Book{
		ID:          id,
		Title:       title,
		Author:      author,
		Price:       float32(price),
		Description: description,
		ImagePath:   filePath, // empty if no new image
		Category:    category,
		IsStock:     is_stock,
	})

	if err8 != nil {
		if err8 == sql.ErrNoRows {
			util.SendError(w, "Book not found", http.StatusNotFound)
			return
		}
		util.SendError(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	// Return updated book
	util.SendData(w, updatedBook, http.StatusOK)
}
