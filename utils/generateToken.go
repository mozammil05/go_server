package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

// Load the environment variables from the .env file
func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

// GenerateToken generates a JWT token based on the given email and role.
func GenerateToken(email string, role string) (string, error) {
	// Get the JWT secret key from the environment variables
	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	// Create claims with email and role
	claims := Claims{}
	claims.Email = email
	claims.Role = role

	// Set token expiration time (e.g., 2 hours)
	expirationTime := time.Now().Add(2 * time.Hour)
	claims.ExpiresAt = expirationTime.Unix()

	// Create the JWT token with claims and sign it using the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
