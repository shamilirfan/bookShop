package handlers

import (
	"bookShop/database"
	"bookShop/util"
	"net/http"
	"strconv"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// Path value
	bookID := r.PathValue("id")
	id, err := strconv.Atoi(bookID)

	// Error handling
	if err != nil {
		http.Error(w, "Please give a valid id", http.StatusBadRequest)
		return
	}

	database.Delete(id)                                     // Call delete function
	util.SendData(w, "Successfully deleted", http.StatusOK) // Call sendData function
}
