// routes/auth.go

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suriya1776/hinata/usermangement-service/api"
)

func SetupAuthRoutes(router *gin.Engine) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/login", api.LoginHandler)
		authGroup.POST("/register", api.RegisterHandler)
		authGroup.POST("/resetpassword", api.ResetPasswordHandler)
		// Other authentication routes...
	}
}
