package middleware

import (
	// ...

	"github.com/gin-gonic/gin"
)

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// email, existsEmail := c.Get("email")
		// role, existsRole := c.Get("role")

		// // Check if both email and role are present in the context
		// if !existsEmail || !existsRole {
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		// 	c.Abort()
		// 	return
		// }

		// // Check if the user has the "admin" role and the correct email
		// if role.(string) != "admin" || email.(string) != "expectedEmail" {
		// 	c.JSON(http.StatusForbidden, gin.H{"error": "Access denied. User is not authorized or has an invalid email"})
		// 	c.Abort()
		// 	return
		// }

		c.Next()
	}
}
