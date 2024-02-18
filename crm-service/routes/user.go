// routes/user_routes.go

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suriya1776/hinata/crm-service/api"
)

func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	userGroup.Use(api.AuthMiddleware()) // Apply authentication middleware
	{
		userGroup.GET("/profile", api.UserProfileHandler)
		// Other user-related routes...

		// Add the following route for updating the secret key
		userGroup.PUT("/update-secret", api.AdminMiddleware(), api.UpdateSecretHandler)
	}
}
