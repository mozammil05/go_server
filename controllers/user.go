package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserProfile(c *gin.Context) {
	// Implement user profile retrieval logic here
	responseMessage := "User profile retrieved successfully"

	// Send a JSON response to the client
	c.JSON(http.StatusOK, gin.H{"message": responseMessage})
}

func UpdateUserProfile(c *gin.Context) {
	// Implement user profile update logic here
	fmt.Printf("hello")
}
