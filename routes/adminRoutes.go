package routes

import (
	"my-auth-app/controllers"
	"my-auth-app/middleware"
	"my-auth-app/utils"

	"github.com/gin-gonic/gin"
)

func SetupAdminRoutes(r *gin.RouterGroup, db *utils.Database) {
	adminRoutes := r.Group("/admin")
	adminRoutes.Use(middleware.AdminMiddleware())
	{
		adminRoutes.POST("/create-profile", func(c *gin.Context) {
			controllers.CreateUserProfile(c,db)
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
}
