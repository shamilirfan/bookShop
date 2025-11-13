package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/joho/godotenv"
)

// Struct define for configaration
type Configaration struct {
	Version      string
	ServiceName  string
	HttpPort     int64
	JwtSecretKey string
	Cloudinary   *cloudinary.Cloudinary
}

// Configaration type variable define
var config *Configaration

// Configaration loading function
func loadConfig() {
	err_1 := godotenv.Load()
	version := os.Getenv("VERSION")
	serviceName := os.Getenv("SERVICE_NAME")
	httpPort := os.Getenv("HTTP_PORT")
	port, err_2 := strconv.ParseInt(httpPort, 10, 64)
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	cloudinary, err_3 := cloudinary.NewFromURL(os.Getenv("CLOUDINARY_URL"))

	if err_1 != nil {
		log.Fatal("❌ Failed to load the env variables", err_1)
	}

	if version == "" {
		fmt.Println("Version is required")
		return
	}

	if serviceName == "" {
		fmt.Println("Service Name is required")
		return
	}

	if httpPort == "" {
		fmt.Println("Port number is required")
		return
	}

	if err_2 != nil {
		fmt.Println("Port must be number")
		return
	}

	if err_3 != nil {
		fmt.Println("Cloudinary is not connected")
	} else {
		fmt.Println("✅ Cloudinary is connected")
	}

	if jwtSecretKey == "" {
		fmt.Println("Jwt secret key is required")
		return
	}

	config = &Configaration{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     port,
		JwtSecretKey: jwtSecretKey,
		Cloudinary:   cloudinary,
	}
}

// Get loadConfig function
func GetConfig() *Configaration {
	if config == nil {
		loadConfig()
	}
	return config
}
