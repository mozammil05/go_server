package middleware

import (
	// ...
	"my-auth-app/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the JWT secret key from environment variables
		jwtSecret := os.Getenv("JWT_SECRET")

		// Get the authorization header
		authHeader := c.Request.Header.Get("Authorization")

		// Check if the header is empty
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		// Split the header to get the token part
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization format"})
			c.Abort()
			return
		}

		// Get the token string
		tokenString := parts[1]

		// Parse and validate the token using the secret key
		token, err := jwt.ParseWithClaims(tokenString, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Check if the token is valid
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is not valid"})
			c.Abort()
			return
		}

		// Get the claims from the token
		claims, ok := token.Claims.(*utils.Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		// Set user information in the Gin context
		c.Set("user", claims)
		// Set user information in the Gin context for use in route handlers
		c.Next()
	}
}

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		// Check if the user has the "user" role
		claims, ok := userClaims.(*utils.Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		if claims.Role != "user" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. User is not authorized"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// AdminMiddleware middleware for admin role
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		// Check if the user has the "admin" role
		claims, ok := userClaims.(*utils.Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		if claims.Role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. User is not an admin"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// SuperAdminMiddleware middleware for superadmin role
func SuperAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userClaims, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		// Check if the user has the "superadmin" role
		claims, ok := userClaims.(*utils.Claims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		if claims.Role != "superadmin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. User is not a superadmin"})
			c.Abort()
			return
		}

		c.Next()
	}
}
