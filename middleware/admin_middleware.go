package middleware

import (
	// ...
	"my-auth-app/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
