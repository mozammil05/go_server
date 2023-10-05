package routes

import (
	"my-auth-app/controllers"
	"my-auth-app/middleware"
	"my-auth-app/utils"

	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(db *utils.Database, jwtSecret string) *gin.Engine {
	r := gin.Default()

	// Initialize JWT
	utils.InitJWT(jwtSecret)

	// Create a route group with a custom prefix
	apiV1 := r.Group("/api/v1")
	{
		// Group routes under "public" prefix within /api/v1
		public := apiV1.Group("/public")
		{
			public.POST("/signup", func(c *gin.Context) {
				// Call the Signup function from controllers package
				controllers.Signup(c, db)
			})
			public.POST("/login", func(c *gin.Context) {
				// Call the Login function from controllers package and pass the db instance
				controllers.Login(c, db)
			})
			public.POST("/resetpassword", controllers.ResetPassword)
		}

		// Group routes under "private" prefix within /api/v1 with authentication middleware
		private := apiV1.Group("/private")
		private.Use(middleware.AuthMiddleware())
		{
			private.POST("/create-profile", func(c *gin.Context) {
				// Check if the authenticated user has the "user" role
				user := c.MustGet("user").(utils.User)
				if user.Role != "user" {
					c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
					return
				}
				controllers.CreateUserProfile(c)
			})

			private.GET("/get-profile", func(c *gin.Context) {
				// Check if the authenticated user has the "admin" or "superuser" role
				user := c.MustGet("user").(utils.User)
				if user.Role != "admin" && user.Role != "superuser" {
					c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
					return
				}
				controllers.GetAllUsers(c, db)
			})

			// Route to update user profile
			private.PUT("/update-profile", func(c *gin.Context) {
				// Check if the authenticated user has the "user" role
				user := c.MustGet("user").(utils.User)
				if user.Role != "user" {
					c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
					return
				}
				controllers.UpdateUserProfile(c, db)
			})

			private.POST("/changepassword", func(c *gin.Context) {
				controllers.ChangePassword(c, db)
			})
		}
	}

	return r
}
