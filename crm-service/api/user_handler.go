// api/handlers.go

package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserProfileHandler(c *gin.Context) {
	// Extract user information from the token
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
		return
	}

	role, exists := c.Get("role")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing token"})
		return
	}

	// Customize this part based on your user data retrieval logic
	userProfile := map[string]interface{}{
		"username": username,
		"role":     role,
		// Add other user details as needed
	}

	c.JSON(http.StatusOK, gin.H{"userProfile": userProfile})
}
