package routes

import (
	"my-auth-app/controllers"
	"my-auth-app/middleware"
	"my-auth-app/utils"

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

		// Routes for regular users
		userRoutes := private.Group("/user")
		userRoutes.Use(middleware.UserMiddleware())
		{
			userRoutes.POST("/create-profile", func(c *gin.Context) {
				controllers.CreateUserProfile(c)
			})
			userRoutes.GET("/get", func(c *gin.Context) {
				controllers.GetAllUsers(c, db)
			})

			userRoutes.PUT("/update-profile", func(c *gin.Context) {
				controllers.UpdateUserProfile(c, db)
			})
			userRoutes.POST("/changepassword", func(c *gin.Context) {
				controllers.ChangePassword(c, db)
			})
		}

		// Routes for admins
		adminRoutes := private.Group("/admin")
		adminRoutes.Use(middleware.AdminMiddleware())
		{
			adminRoutes.POST("/create-profile", func(c *gin.Context) {
				controllers.CreateUserProfile(c)
			})
			adminRoutes.GET("/get-profile", func(c *gin.Context) {
				controllers.GetAllUsers(c, db)
			})

			adminRoutes.PUT("/update-profile", func(c *gin.Context) {
				controllers.UpdateUserProfile(c, db)
			})
			adminRoutes.POST("/changepassword", func(c *gin.Context) {
				controllers.ChangePassword(c, db)
			})
		}

		// Routes for superadmins
		superadminRoutes := private.Group("/superadmin")
		superadminRoutes.Use(middleware.SuperAdminMiddleware())
		{
			superadminRoutes.POST("/create-profile", func(c *gin.Context) {
				controllers.CreateUserProfile(c)
			})
			superadminRoutes.GET("/get-profile", func(c *gin.Context) {
				controllers.GetAllUsers(c, db)
			})

			superadminRoutes.PUT("/update-profile", func(c *gin.Context) {
				controllers.UpdateUserProfile(c, db)
			})
			superadminRoutes.POST("/changepassword", func(c *gin.Context) {
				controllers.ChangePassword(c, db)
			})
		}
	}

	return r
}
