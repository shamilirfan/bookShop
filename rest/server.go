package rest

import (
	"bookShop/config"
	"bookShop/rest/middlewares"
	"fmt"
	"net/http"
	"os"
)

func Server() {
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
	Routes(mux)

	// Listening server
	fmt.Println("Server is running at http://localhost" + port)
	err := http.ListenAndServe(port, mux)

	// Handle Listening server error
	if err != nil {
		fmt.Println("Something went wrong", err)
		os.Exit(1)
	}
}
