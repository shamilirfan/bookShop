package cmd

import (
	"bookShop/config"
	"bookShop/infrastructure/db"
	"bookShop/repo/book"
	usersRep "bookShop/repo/users"
	"bookShop/rest"
	bookHandler "bookShop/rest/handlers/book"
	"bookShop/rest/handlers/users"
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
	usersRepo := usersRep.NewUsersRepo(dbConfigaretion)

	server := rest.NewServer(
		config,
		bookHandler.NewHandler(bookRepo),
		users.NewHandler(usersRepo),
	)

	server.Start()
}
