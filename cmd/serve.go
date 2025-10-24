package cmd

import (
	"bookShop/config"
	"bookShop/rest"
	"bookShop/rest/handlers/book"
)

func Serve() {
	config := config.GetConfig()
	server := rest.NewServer(config, book.NewHandler())
	server.Start()
}
