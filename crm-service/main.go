package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/suriya1776/hinata/crm-service/database"
	"github.com/suriya1776/hinata/crm-service/routes"
)

func main() {
	router := gin.Default()

	// Initialize the database
	err := database.InitializeDB()
	if err != nil {
		log.Fatal("Failed to initialize the database:", err)
	}
	// Register API routes
	routes.SetupAuthRoutes(router)
	routes.SetupUserRoutes(router)

	// Run the server
	router.Run(":8080")
}
