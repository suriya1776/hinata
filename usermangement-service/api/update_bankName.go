// api/bank_handler.go

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suriya1776/hinata/usermangement-service/database"
)

func UpdateBankNameHandler(c *gin.Context) {
	// Check if the user has admin role

	var updateRequest database.UpdateBankNameRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the bank name
	if !database.IsValidBankName(updateRequest.NewBankName) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bank name format. Special characters are allowed."})
		return
	}

	err := database.UpdateBankName(updateRequest.NewBankName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bank name"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bank name updated successfully"})
}
