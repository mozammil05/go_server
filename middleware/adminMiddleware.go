package middleware

import (
	// ...

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if the user is properly authenticated and has the correct role
		email, existsEmail := c.Get("email")
		role, existsRole := c.Get("role")

		fmt.Printf("Role: %+v, Email: %+v\n", role, email)

		if !existsEmail || !existsRole {
			fmt.Println("User not authenticated")
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		if role.(string) != "admin" {
			fmt.Println("Access denied. User is not authorized as an admin")
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. User is not authorized as an admin"})
			c.Abort()
			return
		}

		c.Next()
	}
}
