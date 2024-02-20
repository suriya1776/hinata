package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/suriya1776/hinata/usermangement-service/models"
	"go.mongodb.org/mongo-driver/bson"
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
	user.LastPasswordUpdate = time.Now()
	// Hash the password before inserting
	hashedPassword, err := HashPassword(user.Password)
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

// UpdateUser updates the user password and last password update time in the database
func UpdateUser(user models.BankUser) error {

	fmt.Println("UpdateUser:", user)
	// Assuming you have a MongoDB collection named "users"
	// filter := bson.M{"_id": user.ID}

	filter := bson.M{"_id": user.ID}
	update := bson.M{
		"$set": bson.M{
			"password":           user.Password,
			"lastPasswordUpdate": time.Now(),
		},
	}

	// UpdateOne performs an atomic update
	result, err := usersCollection.UpdateOne(context.TODO(), filter, update)
	fmt.Println("error", err)
	if err != nil {

		return err
	}

	// Check if any documents were matched and modified
	if result.ModifiedCount == 0 {
		return errors.New("no documents were updated")
	}

	// Debug prints
	fmt.Println("UpdateUser Query:", result)

	return nil
}
