package routes

import (
	"my-auth-app/controllers"
	"my-auth-app/middleware"
	"my-auth-app/utils"

	"github.com/gin-gonic/gin"
)

func SetupSuperAdminRoutes(r *gin.RouterGroup, db *utils.Database) {
	superadminRoutes := r.Group("/superadmin")
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
