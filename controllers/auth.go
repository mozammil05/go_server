package controllers

import (
	"context"
	"fmt"
	"my-auth-app/models"
	"my-auth-app/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context, db *utils.Database) {
	var userInput models.User

	// Bind user input from the request body
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the email is already taken
	existingUser, err := db.UserCollection.FindOne(context.TODO(), bson.M{"email": userInput.Email}).DecodeBytes()
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}
	if existingUser != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
		return
	}

	// Hash the user's password
	hashedPassword, err := utils.HashPassword(userInput.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	// Create a new user document
	newUser := models.User{
		Email:    userInput.Email,
		Username: userInput.Username,
		Password: hashedPassword,
	}

	// Insert the user into the database
	_, err = db.UserCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "error": false, "data": newUser})

}

// login

type UserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

func Login(c *gin.Context, db *utils.Database) {
	var userInput models.User

	// Bind user input from the request body
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user by email in the database
	filter := bson.M{"email": userInput.Email}
	existingUser := &models.User{}
	err := db.UserCollection.FindOne(context.TODO(), filter).Decode(existingUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Compare the hashed password from the database with the input password
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(userInput.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate a JWT token
	token, err := utils.CreateToken(existingUser.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	istLocation, _ := time.LoadLocation("Asia/Kolkata")
	expirationTime := time.Now().In(istLocation).Add(2 * time.Hour)

	// Format the expiration time in AM/PM format
	expirationTimeFormatted := expirationTime.Format("2006-01-02 03:04 PM MST")

	// Set the token in the response header
	c.Header("Authorization", "Bearer "+token)

	// Create a UserResponse object with the fields you want to include in the response
	userResponse := UserResponse{
		Email:    existingUser.Email,
		Username: existingUser.Username,
	}

	// Return the response with the formatted IST expiration time
	c.JSON(http.StatusOK, gin.H{
		"message":          "Login successful",
		"token_expires_at": expirationTimeFormatted,
		"token":            token,
		"data":             userResponse,
	})
}

func ChangePassword(c *gin.Context, db *utils.Database) {
	var changePasswordInput models.ChangePasswordInput

	// Bind change password input from the request body
	if err := c.ShouldBindJSON(&changePasswordInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the old password
	// Query the user by email
	filter := bson.M{"email": changePasswordInput.Email}
	existingUser := &models.User{}
	err := db.UserCollection.FindOne(context.TODO(), filter).Decode(existingUser)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Debugging: Print existing and input passwords
	fmt.Println("Existing Password:", existingUser.Password)
	fmt.Println("Input Old Password:", changePasswordInput.OldPassword)

	// Compare the old password with the input old password
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(changePasswordInput.OldPassword)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid old password"})
		return
	}

	// Hash the new password
	hashedNewPassword, err := utils.HashPassword(changePasswordInput.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash new password"})
		return
	}

	// Update the user's password in the database
	update := bson.M{
		"$set": bson.M{
			"password": hashedNewPassword,
		},
	}
	filter = bson.M{"email": changePasswordInput.Email}
	_, err = db.UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password changed successfully"})
}

func ResetPassword(c *gin.Context) {
	// Implement reset password logic here
}
