package main

import (
	"github.com/suriya1776/hinata/crm-service/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Register API routes
	api.RegisterRoutes(router)

	// Run the server
	router.Run(":8080")
}
