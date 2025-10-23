package handlers

import (
	"bookShop/database"
	"bookShop/util"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	util.SendData(w, database.BookList(), http.StatusOK)
}
