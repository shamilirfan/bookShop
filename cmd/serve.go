package cmd

import (
	"bookShop/config"
	"bookShop/repo"
	"bookShop/rest"
	"bookShop/rest/handlers/book"
)

func Serve() {
	config := config.GetConfig()

	bookRepo := repo.NewBookRepo()
	server := rest.NewServer(config, book.NewHandler(bookRepo))

	server.Start()
}
