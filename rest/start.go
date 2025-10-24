package rest

import (
	"bookShop/config"
	"bookShop/rest/handlers/book"
	"bookShop/rest/middlewares"
	"fmt"
	"net/http"
	"os"
)

type Server struct {
	config      *config.Configaration
	bookHandler *book.Handler
}

func NewServer(
	config *config.Configaration,
	bookHandler *book.Handler,
) *Server {
	return &Server{
		config:      config,
		bookHandler: bookHandler,
	}
}

func (server *Server) Start() {
	// Port
	var port string = ":" + fmt.Sprintf("%d", config.GetConfig().HttpPort)

	// Call ChainMiddleware function and pass argument
	middlewares.ChainMiddleware(
		middlewares.CorsWithPreflight,
		middlewares.Logger, // Start will from here
	)

	// Create router
	mux := http.NewServeMux()

	// Call routes function and pass mux as a argument
	// Routes(mux)
	// book.NewHandler().RregisterRoutes(mux)
	server.bookHandler.RregisterRoutes(mux)

	// Listening server
	fmt.Println("Server is running at http://localhost" + port)
	err := http.ListenAndServe(port, mux)

	// Handle Listening server error
	if err != nil {
		fmt.Println("Something went wrong", err)
		os.Exit(1)
	}
}
