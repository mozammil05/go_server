package controllers

import (
	"context"
	"log"
	"my-auth-app/models"
	"my-auth-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProfile handles the creation of a user profile.
func CreateUserProducts(c *gin.Context, db *utils.Database) {
	// Define a variable to hold the user data
	var userInput models.Product

	// Bind the JSON request body to the userInput variable
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log the received user data
	log.Printf("Received product data: %+v", userInput)

	// Create a new user object with the hashed password
	newUser := models.Product{
		ID:          userInput.ID,
		Name:        userInput.Name,
		Description: userInput.Description,
		Price:       userInput.Price,
		Category:    userInput.Category,
	}

	// Insert the user into the database
	insertResult, err := db.ProductCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Failed to create user",
			"status": http.StatusInternalServerError,
		})
		log.Printf("Error inserting user into the database: %v", err)
		return
	}

	// Send a JSON response to the client.
	c.JSON(http.StatusCreated, gin.H{
		"message": "User product created successfully",
		"status":  http.StatusCreated,
		"error":   false,
		"data":    insertResult.InsertedID, // Assuming you want to return the inserted document ID
	})
}
