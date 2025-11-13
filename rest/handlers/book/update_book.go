package book

import (
	"bookShop/repo/book"
	"bookShop/util"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

// Helper function to extract Cloudinary PublicID
func cloudinaryPublicID(imageURL string) string {
	parts := strings.Split(imageURL, "/upload/")
	if len(parts) < 2 {
		return ""
	}
	publicPath := parts[1]
	publicPath = strings.SplitN(publicPath, ".", 2)[0]
	return publicPath
}

func (h *Handler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	err1 := r.ParseMultipartForm(50 << 20)
	bookID := r.PathValue("id")
	id, err2 := strconv.Atoi(bookID)
	title := r.FormValue("title")
	author := r.FormValue("author")
	priceStr := r.FormValue("price")
	price, err3 := strconv.ParseFloat(priceStr, 64)
	description := r.FormValue("description")
	category := r.FormValue("category")
	brand := r.FormValue("brand")
	isStockStr := r.FormValue("is_stock")
	is_stock, err5 := strconv.ParseBool(isStockStr)
	files := r.MultipartForm.File["image_path"]

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

	// files upload limitation
	if len(files) > 4 {
		http.Error(w, "Maximum 4 images allowed", http.StatusBadRequest)
		return
	}

	// Delete old images from Cloudinary
	var booK book.Book
	for _, imgURL := range booK.ImagePath {
		publicID := cloudinaryPublicID(imgURL)
		if publicID == "" {
			continue
		}
		_, err := h.cld.Upload.Destroy(r.Context(), uploader.DestroyParams{
			PublicID: publicID,
		})
		if err != nil {
			fmt.Println("Failed to delete old image:", imgURL, err)
		} else {
			fmt.Println("Deleted old image:", imgURL)
		}
	}

	// Upload new images
	var imageUrls []string
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			http.Error(w, "Error reading file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		uploadResult, err := h.cld.Upload.Upload(r.Context(), file, uploader.UploadParams{
			Folder: "books",
		})
		if err != nil {
			http.Error(w, "Cloudinary upload failed", http.StatusInternalServerError)
			return
		}

		imageUrls = append(imageUrls, uploadResult.SecureURL)
	}

	// Update book
	b := book.Book{
		ID:          id,
		Title:       title,
		Author:      author,
		Price:       float32(price),
		Description: description,
		ImagePath:   imageUrls,
		Category:    category,
		Brand:       brand,
		IsStock:     is_stock,
	}
	updatedBook, err8 := h.bookRepo.Update(b)

	if err8 != nil {
		util.SendError(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	// Return updated book
	util.SendData(w, updatedBook, http.StatusOK)
}
