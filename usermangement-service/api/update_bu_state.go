package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/suriya1776/hinata/usermangement-service/database"
	"github.com/suriya1776/hinata/usermangement-service/models"
)

func AddRoleHandler(c *gin.Context) {
	var newRole models.Role
	if err := c.ShouldBindJSON(&newRole); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the role name is "admin"
	roleName := strings.ToUpper(newRole.Name)
	if roleName == "ADMIN" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot create role with the name 'admin'"})
		c.Abort()
		return
	}

	err := database.InsertRole(newRole)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Role added successfully", "role": newRole.Name})
}

// GrantRolesHandler allows admin users to grant roles to users
func GrantRolesHandler(c *gin.Context) {
	// Apply AdminMiddleware
	AdminMiddleware()(c)

	var grantRolesRequest models.GrantRolesRequest
	if err := c.ShouldBindJSON(&grantRolesRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate that the roles being granted are valid roles in the database
	for _, roleName := range grantRolesRequest.Roles {
		if !database.IsValidRole(roleName) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid role: " + roleName})
			return
		}
	}

	// Grant roles to the user
	err := database.GrantRoles(grantRolesRequest.Username, grantRolesRequest.Roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User does not exist"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Roles granted successfully"})
}
