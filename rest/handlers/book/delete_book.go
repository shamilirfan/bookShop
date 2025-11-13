package book

import (
	"bookShop/repo/book"
	"bookShop/util"
	"fmt"
	"net/http"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Path value
	bookID := r.PathValue("id")
	id, err := strconv.Atoi(bookID)

	// Delete images from Cloudinary
	var booK book.Book
	for _, imgURL := range booK.ImagePath {
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
