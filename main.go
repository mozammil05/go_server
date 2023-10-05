package main

import (
	"fmt"
	"log"
	"my-auth-app/routes"
	"my-auth-app/utils"
	"os"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get the JWT secret key from the environment variables
	jwtSecret := os.Getenv("JWT")

	// Initialize MongoDB connection
	db := utils.InitDB()
	defer db.Disconnect()

	// Check if the MongoDB connection was successful
	if db.Client != nil {
		fmt.Println("MongoDB connected successfully")
	}

	// Get the port number from the environment variables
	port := os.Getenv("PORT")

	// Use a default port if the environment variable is not set
	if port == "" {
		port = "8080" // Default port
	}

	// Create a new router
	router := routes.NewRouter(db, jwtSecret)

	// Start your server on the specified port
	router.Run(":" + port)
}
