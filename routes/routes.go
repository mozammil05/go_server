package routes

import (
	"my-auth-app/controllers"
	"my-auth-app/middleware"
	"my-auth-app/utils"

	"github.com/gin-gonic/gin"
)

func NewRouter(db *utils.Database, jwtSecret string) *gin.Engine {
	r := gin.Default()

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

		// Include admin, user, and superadmin routes from separate router files
		SetupUserRoutes(private, db)
		SetupAdminRoutes(private, db)
		SetupSuperAdminRoutes(private, db)
	}

	return r
}
