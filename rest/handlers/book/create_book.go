package book

import (
	"bookShop/repo/book"
	"bookShop/util"
	"net/http"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func (h *Handler) CreateBook(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(50 << 20)
	if err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	title := r.FormValue("title")
	author := r.FormValue("author")
	priceStr := r.FormValue("price")
	price, err := strconv.ParseFloat(priceStr, 32)
	if err != nil {
		util.SendError(w, "Invalid price format", http.StatusBadRequest)
		return
	}

	description := r.FormValue("description")
	category := r.FormValue("category")
	brand := r.FormValue("brand")
	is_stock := r.FormValue("is_stock")
	stock, err := strconv.ParseBool(is_stock)
	if err != nil {
		util.SendError(w, "Invalid stock value", http.StatusBadRequest)
		return
	}

	files := r.MultipartForm.File["image_path"]
	if len(files) > 4 {
		http.Error(w, "You can upload a maximum of 4 images", http.StatusBadRequest)
		return
	}

	var imageUrls []string

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			util.SendError(w, "Error reading file", http.StatusInternalServerError)
			return
		}

		uploadResult, err := h.cld.Upload.Upload(r.Context(), file, uploader.UploadParams{
			Folder: "books",
		})
		file.Close() // ✅ লুপের ভিতরেই বন্ধ করে দাও

		if err != nil {
			util.SendError(w, "Cloudinary upload failed", http.StatusInternalServerError)
			return
		}

		imageUrls = append(imageUrls, uploadResult.SecureURL)
	}

	createdBook, err := h.bookRepo.Create(book.Book{
		Title:       title,
		Author:      author,
		Price:       float32(price),
		Description: description,
		ImagePath:   imageUrls,
		Category:    category,
		Brand:       brand,
		IsStock:     stock,
	})

	if err != nil {
		util.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, createdBook, http.StatusCreated)
}
