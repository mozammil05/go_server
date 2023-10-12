package controllers

import (
	"context"
	"fmt"
	"log"
	"my-auth-app/models"
	"my-auth-app/services"
	"my-auth-app/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// CreateProfile handles the creation of a user profile.
func CreateUserProfile(c *gin.Context, db *utils.Database) {
	// Define a variable to hold the user data
	var userInput models.User

	// Bind the JSON request body to the userInput variable
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Log the received user data
	log.Printf("Received user data: %+v", userInput)

	// Create a new user object with the hashed password
	newUser := models.User{
		Email:    userInput.Email,
		Username: userInput.Username,
		Role:     userInput.Role,
		Created:  userInput.Created,
		Updated:  userInput.Updated,
	}

	// Insert the user into the database
	insertResult, err := db.UserCollection.InsertOne(context.TODO(), newUser)
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
		"message": "User profile created successfully",
		"status":  http.StatusCreated,
		"error":   false,
		"data":    insertResult.InsertedID, // Assuming you want to return the inserted document ID
	})
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
		c.JSON(http.StatusNotFound, gin.H{"error": "No users found", "status": http.StatusNotFound})
		return
	}

	// Create a response object without the "password" field for each user
	var response []utils.UserProfileResponse

	for _, user := range users {
		response = append(response, utils.UserProfileResponse{
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

type UpdateResponse struct {
	Username     string `json:"username"`
	ProfileImage string `json:"profileImage"`
}

func UpdateUserProfile(c *gin.Context, db *utils.Database) {
	// Get the user's email from the token claims in the context
	userEmail, exists := c.Get("email")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Parse form data including username and files
	if err := c.Request.ParseMultipartForm(10 << 20); err != nil { // 10 MB limit
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse form data"})
		return
	}

	// Handle file upload
	filePaths, err := services.HandleUpload(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload the image"})
		return
	}

	// Assuming only one file is uploaded, you can change this logic if multiple files are uploaded
	if len(filePaths) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No image file provided"})
		return
	}

	filePath := filePaths[0] // Using the first file, you can loop through if there are multiple files

	// Get the username from the form
	username := c.PostForm("username")
	// Calculate the relative file path
	relativeFilePath := "destination/" + filepath.Base(filePath)

	// Update the user's profile in the database
	filter := bson.M{"email": userEmail}
	update := bson.M{
		"$set": bson.M{
			"username":     username,
			"profileImage": relativeFilePath,
			// Add other fields you want to update here
		},
	}

	// Perform the database update
	_, updateErr := db.UserCollection.UpdateOne(context.TODO(), filter, update)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user profile in the database"})
		return
	}

	// Set response data
	responseData := UpdateResponse{
		Username:     username,
		ProfileImage: relativeFilePath, // Use the relative file path
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User profile updated successfully",
		"status":  http.StatusOK,
		"error":   false,
		"data":    responseData,
	})
}

// GetProfile retrieves the user's profile.
func GetProfile(c *gin.Context, db *utils.Database) {
	// Get the user's email from the token claims in the context
	userEmail, exists := c.Get("email")
	fmt.Printf("Email:  %s\n", userEmail)

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Define a filter to find the user by email
	filter := bson.M{"email": userEmail}

	// Find the user in the database
	var user models.User
	err := db.UserCollection.FindOne(context.TODO(), filter).Decode(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user profile"})
		return
	}

	// Create a response object without the "password" field
	userProfile := utils.UserProfileResponse{
		Email:    user.Email,
		Username: user.Username,
		Role:     user.Role,
		// Add other profile fields here
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"error":   false,
		"data":    userProfile,
		"message": "User profile retrieved successfully",
	})
}
