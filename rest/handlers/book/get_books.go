package book

import (
	"bookShop/database"
	"bookShop/util"
	"net/http"
)

func (h *Handler) GetBooks(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.BookList(), http.StatusOK)
}
