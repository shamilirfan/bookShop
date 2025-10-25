package cmd

import (
	"bookShop/config"
	"bookShop/infrastructure/db"
	"bookShop/repo/book"
	"bookShop/rest"
	bookHandler "bookShop/rest/handlers/book"
	"fmt"
	"os"
)

func Serve() {
	config := config.GetConfig()
	dbConfigaretion, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bookRepo := book.NewBookRepo(dbConfigaretion)
	server := rest.NewServer(config, bookHandler.NewHandler(bookRepo))

	server.Start()
}
