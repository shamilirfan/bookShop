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

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Path value
	bookID := r.PathValue("id")
	id, err := strconv.Atoi(bookID)

	// Delete images from Cloudinary
	var book book.Book
	for _, imgURL := range book.ImagePath {
		publicID := cloudinaryPublicID(imgURL)
		if publicID == "" {
			fmt.Println("❌ Invalid Cloudinary URL:", imgURL)
			continue
		}

		_, err := h.cld.Upload.Destroy(r.Context(), uploader.DestroyParams{
			PublicID: publicID,
		})
		if err != nil {
			fmt.Println("❌ Failed to delete from Cloudinary:", imgURL, err)
		} else {
			fmt.Println("✅ Deleted from Cloudinary:", imgURL)
		}
	}

	// Error handling
	if err != nil {
		util.SendError(w, "Please give a valid id", http.StatusBadRequest)
		return
	}

	h.bookRepo.Delete(id)                                   // Call delete function
	util.SendData(w, "Successfully deleted", http.StatusOK) // Call sendData function
}
