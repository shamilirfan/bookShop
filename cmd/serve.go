package cmd

import (
	"bookShop/rest"
	"bookShop/rest/handlers/book"
)

func Serve() {
	server := rest.NewServer(book.NewHandler())
	server.Start()
}
