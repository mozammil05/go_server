package controllers

import (
	"context"
	"my-auth-app/models"
	"my-auth-app/utils"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
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
	newUser := models.User{
		Email:    userInput.Email,
		Username: userInput.Username,
		Password: hashedPassword,
		Role:     userInput.Role,
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
type UserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func Login(c *gin.Context, db *utils.Database) {
	var userInput models.User

	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

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

	// Determine the user's role (user, admin, superadmin)
	userRole := existingUser.Role // You need to define the role property in your User model

	// Your JWT secret key
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	// Create claims with email and role
	claims := utils.Claims{}
	claims.Email = userInput.Email
	claims.Role = userRole

	// Set token expiration time (e.g., 2 hours)
	istLocation, _ := time.LoadLocation("Asia/Kolkata")
	expirationTime := time.Now().In(istLocation).Add(2 * time.Hour)
	claims.ExpiresAt = expirationTime.Unix()

	// Format the expiration time in AM/PM format
	expirationTimeFormatted := expirationTime.Format("2006-01-02 03:04 PM MST")

	// Create the JWT token with claims and sign it using the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// Set the token in the response header
	c.Header("Authorization", "Bearer "+tokenString)

	// Create a UserResponse object with the fields you want to include in the response
	userResponse := UserResponse{
		Email:    existingUser.Email,
		Username: existingUser.Username,
		Role:     userRole,
	}

	// Return the response
	c.JSON(http.StatusOK, gin.H{
		"message":          "Login successful",
		"token":            tokenString,
		"token_expires_at": expirationTimeFormatted,
		"data":             userResponse,
		"status":           http.StatusOK,
		"error":            false,
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

func ResetPassword(c *gin.Context) {
	// Implement reset password logic here
}
