package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

func ProjectConfig() {
	err_1 := godotenv.Load()
	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	httpPort := os.Getenv("HTTP_PORT")
	port, err_2 := strconv.ParseInt(httpPort, 10, 64)
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	if err_1 != nil {
		fmt.Println("Failed to load the env variables", err_1)
	}

	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	if serviceName == "" {
		fmt.Println("Service Name is required")
		os.Exit(1)
	}

	if httpPort == "" {
		fmt.Println("Port number is required")
		os.Exit(1)
	}

	if err_2 != nil {
		fmt.Println("Port must be number")
		os.Exit(1)
	}

	if jwtSecretKey == "" {
		fmt.Println("Jwt secret key is required")
		os.Exit(1)
	}

	config = &Configaration{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     port,
		JwtSecretKey: jwtSecretKey,
	}
}
