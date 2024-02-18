package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suriya1776/hinata/crm-service/database"
	"github.com/suriya1776/hinata/crm-service/models"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/register", RegisterHandler)
	router.POST("/login", LoginHandler)
}

// Registering the new user
// Does not include the master user

func RegisterHandler(c *gin.Context) {
	var newUser models.BankUser
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the username is valid (only text and numbers allowed, no special characters)
	if !database.IsValidUsername(newUser.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username format. Only text and numbers are allowed."})
		return
	}
	// Check if the username is unique
	if err := database.IsUsernameUnique(newUser.Username); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists. Please choose a different username."})
		return
	}

	err := database.RegisterUser(c, newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user": newUser.Username})
}
