// routes/bank_routes.go

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/suriya1776/hinata/usermangement-service/api"
)

func SetupBankRoutes(router *gin.Engine) {
	bankGroup := router.Group("/bank")
	bankGroup.Use(api.AuthMiddleware()) // Apply authentication middleware
	{
		// Only admin users can update the bank name
		bankGroup.PUT("/update", api.AdminMiddleware(), api.UpdateBankNameHandler)
		bankGroup.POST("/addroles", api.AdminMiddleware(), api.AddRoleHandler)
		bankGroup.POST("/grantroles", api.AdminMiddleware(), api.GrantRolesHandler)

	}
}
