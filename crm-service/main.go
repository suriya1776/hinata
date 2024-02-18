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
	err := database.InitializeMasterUser()
	if err != nil {
		log.Fatal("Failed to initialize the master user:", err)
	}

	err = database.InitializeBankState()
	if err != nil {
		log.Fatal("Failed to initialize the bank state:", err)
	}

	// Register API routes
	routes.SetupAuthRoutes(router)
	routes.SetupUserRoutes(router)

	// Run the server
	router.Run(":8080")
}
