package controllers

import (
	"context"
	"fmt"
	"my-auth-app/models"
	"my-auth-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateProfile handles the creation of a user profile.
func CreateUserProfile(c *gin.Context) {
	// Define a variable to hold the user data
	var userData models.User

	// Bind the JSON request body to the userData variable
	if err := c.ShouldBindJSON(&userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Implement user profile creation logic here.
	// Example: You can create a new user profile by inserting the userData into your database.

	// Debugging: Print the userData to verify that it's received correctly.
	fmt.Println("User data received:", userData)

	// Send a JSON response to the client.
	c.JSON(http.StatusCreated, gin.H{
		"message": "User profile created successfully",
		"status":  http.StatusCreated,
		"error":   false,
	})
}

type UserProfileResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

// GetAllUsers retrieves all users from the database and sends them as a JSON response.
func GetAllUsers(c *gin.Context, db *utils.Database) {
	// Query all users in the database
	cursor, err := db.UserCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users", "status": http.StatusInternalServerError})
		return
	}
	defer cursor.Close(context.TODO())

	var users []models.User

	// Iterate through the cursor and decode each user into the users slice
	for cursor.Next(context.TODO()) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode user data", "status": http.StatusInternalServerError})
			return
		}
		// Append the user to the users slice
		users = append(users, user)
	}

	// Check if there are no users
	if len(users) == 0 {
		if len(users) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "No users found", "status": http.StatusNotFound})
			return
		}
		return
	}

	// Create a response object without the "password" field for each user
	var response []UserProfileResponse
	for _, user := range users {
		response = append(response, UserProfileResponse{
			Email:    user.Email,
			Username: user.Username,
			Role:     user.Role,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"error":   false,
		"message": "Users retrieved successfully",
		"data":    response,
	})

}

// UpdateUserProfile updates a user's profile
func UpdateUserProfile(c *gin.Context, db *utils.Database) {
	// Get the user's email from the request or your authentication method
	var userData models.User

	// Bind the updated user profile from the request body
	var updatedUser models.User
	if err := c.ShouldBindJSON(&updatedUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Define a filter to find the user by email
	filter := bson.M{"email": userData.Email}

	// Create an update document
	update := bson.M{
		"$set": bson.M{
			"username": updatedUser.Username,
			// Add other fields you want to update here
		},
	}

	// Update the user's profile in the database
	_, err := db.UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User profile updated successfully",
		"status":  http.StatusOK,
		"error":   false,
	})
}
