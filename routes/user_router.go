package routes

import (
	"my-auth-app/controllers"
	"my-auth-app/middleware"
	"my-auth-app/utils"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(r *gin.RouterGroup, db *utils.Database) {
	userRoutes := r.Group("/user")
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
}
