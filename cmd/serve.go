package cmd

import (
	"bookShop/config"
	"bookShop/infrastructure/db"
	"bookShop/repo"
	"bookShop/rest"
	"bookShop/rest/handlers/book"
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

	bookRepo := repo.NewBookRepo(dbConfigaretion)
	server := rest.NewServer(config, book.NewHandler(bookRepo))

	server.Start()
}
