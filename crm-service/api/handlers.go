package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suriya1776/hinata/crm-service/database"
	"github.com/suriya1776/hinata/crm-service/models"
)

func RegisterRoutes(router *gin.Engine) {
	// Define your routes here
	router.POST("/register", RegisterHandler)
	// Add more routes as needed
}

func RegisterHandler(c *gin.Context) {
	var newUser models.BankUser
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := database.RegisterUser(c, newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": newUser.Username})
}
