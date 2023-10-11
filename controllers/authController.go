package controllers

import (
	"context"
	"fmt"
	"my-auth-app/models"
	"my-auth-app/services"
	"my-auth-app/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

// signup
func Signup(c *gin.Context, db *utils.Database) {
	var userInput models.User

	customValidator := utils.NewCustomValidator()

	// Bind user input from the request body and validate
	if err := c.ShouldBindJSON(&userInput); err != nil {
		// Handle validation errors
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use the custom validator to further validate the user input
	if err := customValidator.Validate(userInput); err != nil {
		fmt.Println(customValidator)

		// Handle custom validation errors with custom error messages
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the email is already taken
	existingUser, err := db.UserCollection.FindOne(context.TODO(), bson.M{"email": userInput.Email}).DecodeBytes()
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error":  "Email already exists",
			"status": http.StatusConflict})
		return
	}
	if existingUser != nil {
		c.JSON(http.StatusConflict, gin.H{
			"error":  "Email already exists",
			"status": http.StatusConflict,
		})
		return
	}

	// Hash the user's password
	hashedPassword, err := utils.HashPassword(userInput.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Failed to hash password",
			"status": http.StatusInternalServerError,
		})
		return
	}

	// Create a new user document
	newUser := utils.SignResponse{
		Email:    userInput.Email,
		Username: userInput.Username,
		Password: hashedPassword,
		Role:     userInput.Role,
		Created:  userInput.Created,
		Updated:  userInput.Updated,
	}

	// Insert the user into the database
	_, err = db.UserCollection.InsertOne(context.TODO(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":  "Failed to create user",
			"status": http.StatusInternalServerError,
		})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"error":   false,
		"data":    newUser,
		"status":  http.StatusOK,
	})
}

// Login handles user login and generates a JWT token based on the user's role.

func Login(c *gin.Context, db *utils.Database) {
	var userInput models.User

	// Bind user input from the request body
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Find the user by email and role in the database
	filter := bson.M{"email": userInput.Email, "role": userInput.Role}
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

	// Determine the user's role (user, admin, superadmin)
	userRole := existingUser.Role

	// Your JWT secret key (it's better to set it as an environment variable)
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	// Create claims with email and role
	claims := jwt.MapClaims{
		"email": userInput.Email, // Add email claim
		"role":  userRole,        // Add role claim
	}

	// Create the JWT token with claims and sign it using the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set the token in the response header
	c.Header("Authorization", "Bearer "+tokenString)

	// Print email and role for verification
	fmt.Printf("Email: %s, Role: %s\n", userInput.Email, userRole)

	// Create a UserResponse object with the fields you want to include in the response
	userResponse := utils.UserResponse{
		Email:    existingUser.Email,
		Username: existingUser.Username,
		Role:     userRole,
	}

	// Update the user document in the database with the new token information
	update := bson.M{
		"$set": bson.M{
			"is_active":  true,
			"expiration": time.Now().Add(2 * time.Hour), // Set the token expiration time
			"tokens":     tokenString,                   // Replace the existing tokens with the new token string
		},
	}

	_, err = db.UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user document"})
		return
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   tokenString,
		"data":    userResponse,
		"status":  http.StatusOK,
		"error":   false,
	})
}

// func Login(c *gin.Context, db *utils.Database) {
// 	var userInput models.User

// 	if err := godotenv.Load(); err != nil {
// 		panic("Error loading .env file")
// 	}

// 	// Bind user input from the request body
// 	if err := c.ShouldBindJSON(&userInput); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Find the user by email in the database
// 	filter := bson.M{"email": userInput.Email, "role": userInput.Role}
// 	existingUser := &models.User{}
// 	err := db.UserCollection.FindOne(context.TODO(), filter).Decode(existingUser)

// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
// 		return
// 	}

// 	// Compare the hashed password from the database with the input password
// 	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(userInput.Password)); err != nil {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
// 		return
// 	}

// 	// Determine the user's role (user, admin, superadmin)
// 	userRole := existingUser.Role

// 	// Your JWT secret key
// 	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

// 	// Create claims with email and role
// 	claims := utils.Claims{}
// 	claims.Email = userInput.Email
// 	claims.Role = userRole

// 	// Set token expiration time (e.g., 2 hours)
// 	istLocation, _ := time.LoadLocation("Asia/Kolkata")
// 	expirationTime := time.Now().In(istLocation).Add(20 * time.Hour)
// 	claims.ExpiresAt = expirationTime.Unix()

// 	// Format the expiration time in AM/PM format
// 	expirationTimeFormatted := expirationTime.Format("2006-01-02 03:04 PM MST")

// 	// Create the JWT token with claims and sign it using the secret key
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(jwtSecret)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
// 		return
// 	}

// 	// Set the token in the response header
// 	c.Header("Authorization", "Bearer "+tokenString)

// 	// Create a UserResponse object with the fields you want to include in the response
// 	userResponse := utils.UserResponse{
// 		Email:    existingUser.Email,
// 		Username: existingUser.Username,
// 		Role:     userRole,
// 	}

// 	// Return the response
// 	c.JSON(http.StatusOK, gin.H{
// 		"message":          "Login successful",
// 		"token":            tokenString,
// 		"token_expires_at": expirationTimeFormatted,
// 		"data":             userResponse,
// 		"status":           http.StatusOK,
// 		"error":            false,
// 		"is_active":        true,
// 	})

// 	// Update the user document in the database with the new token information
// 	update := bson.M{
// 		"$set": bson.M{
// 			"is_active":  true,
// 			"expiration": expirationTime,
// 			"tokens":     tokenString, // Replace the existing tokens with the new token string
// 		},
// 	}

// 	_, err = db.UserCollection.UpdateOne(context.TODO(), filter, update)
// 	if err != nil {
// 		fmt.Println("Failed to update user document:", err)
// 	}

// }

// change password
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

	c.JSON(http.StatusOK, gin.H{
		"message": "Password changed successfully",
		"status":  http.StatusOK,
		"error":   false,
	})

}

// ResetPassword handles the password reset logic
func ResetPassword(c *gin.Context) {
	// Parse the reset token and new password from the request
	resetToken := c.PostForm("token")
	// newPassword := c.PostForm("new_password")

	// Call the authentication service to reset the password
	// err := services.ResetPassword(resetToken, newPassword)
	// if err != nil {
	//     c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//     return
	// }

	// Retrieve the recipient's email from your database
	userEmail, err := services.GetUserEmailByResetToken(resetToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user email"})
		return
	}

	// Construct the reset link
	resetLink := "https://example.com/reset-password?token=" + resetToken

	// Send the password reset confirmation email
	err = services.SendResetEmail(userEmail, resetLink)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send reset email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}

// // ResetPassword handles the password reset logic
// func ResetPassword(c *gin.Context) {
// 	// Parse the reset token and new password from the request
// 	resetToken := c.PostForm("token")
// 	newPassword := c.PostForm("new_password")

// 	// Call the SendResetEmail function from the emailServices package
// 	err := services.SendResetEmail(resetToken, newPassword)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
// }
