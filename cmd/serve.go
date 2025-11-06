package cmd

import (
	"bookShop/config"
	"bookShop/infrastructure/db"
	"bookShop/repo/book"
	"bookShop/repo/orders"
	usersRep "bookShop/repo/users"
	"bookShop/rest"
	bookHandler "bookShop/rest/handlers/book"
	o "bookShop/rest/handlers/orders"
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
	orderRepo := orders.NewOrderRepo(dbConfigaretion)

	server := rest.NewServer(
		config,
		bookHandler.NewHandler(bookRepo),
		users.NewHandler(usersRepo),
		o.NewHandler(orderRepo),
	)

	server.Start()
}
