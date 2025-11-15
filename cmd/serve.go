package cmd

import (
	"bookShop/config"
	"bookShop/infrastructure/db"
	"bookShop/repo/book"
	"bookShop/repo/cart"
	"bookShop/repo/orders"
	usersRep "bookShop/repo/users"
	"bookShop/rest"
	bookHandler "bookShop/rest/handlers/book"
	c "bookShop/rest/handlers/cart"
	o "bookShop/rest/handlers/orders"
	"bookShop/rest/handlers/users"
	"fmt"
	"os"
)

func Serve() {
	// Config → DB → Repo → Handler → Server

	config := config.GetConfig()
	dbConfigaretion, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	bookRepo := book.NewBookRepo(dbConfigaretion)
	usersRepo := usersRep.NewUsersRepo(dbConfigaretion)
	orderRepo := orders.NewOrderRepo(dbConfigaretion)
	cartRepo := cart.NewCartRepo(dbConfigaretion)

	server := rest.NewServer(
		config,
		bookHandler.NewHandler(bookRepo, config.Cloudinary),
		users.NewHandler(usersRepo),
		o.NewHandler(orderRepo),
		c.NewHandler(cartRepo),
	)

	server.Start()
}
