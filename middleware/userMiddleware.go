package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		email, existsEmail := c.Get("email")
		role, existsRole := c.Get("role")

		fmt.Printf("Role: %+v, Email: %+v\n", role, email)
		if !existsEmail || !existsRole {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
			c.Abort()
			return
		}

		if role.(string) != "user" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. User is not authorized or has an invalid email"})
			c.Abort()
			return
		}

		c.Next()
	}
}
