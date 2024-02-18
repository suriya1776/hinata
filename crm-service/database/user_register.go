package database

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/suriya1776/hinata/crm-service/models"
)

// RegisterUser creates a new user in the "users" collection
func RegisterUser(c *gin.Context, user models.BankUser) error {

	// Use the bank name from the admin user
	//Avoiding creation duplicate bank names
	// Only admin have access to change the very sensitive infos

	// Perform password strength check
	err := isStrongPassword(user.Password)
	if err != nil {
		return err
	}

	// Hash the password before inserting
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Always set the role to "user" for registration
	// in future admin can alter the role of every users
	// Always set the role to "user" for registration
	// In the future, admin can alter the role of every user
	user.Roles = []string{"user"}

	// Insert the user
	_, err = usersCollection.InsertOne(context.TODO(), user)
	return err
}
