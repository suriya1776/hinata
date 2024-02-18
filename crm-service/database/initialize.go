// database/initialize.go
package database

import (
	"context"
	"log"

	"github.com/suriya1776/hinata/crm-service/models"
	"go.mongodb.org/mongo-driver/bson"
)

// InitializeDB inserts a master user into the database
func InitializeDB() error {
	// Check if the master user already exists
	count, err := usersCollection.CountDocuments(context.TODO(), bson.M{"username": "admin"})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Master user already exists. Skipping initialization.")
		return nil
	}

	// Hash the master user's password before inserting
	hashedMasterPassword, err := hashPassword("adminpassword")
	if err != nil {
		return err
	}

	// Create the master user
	masterUser := models.BankUser{
		BankName:    "Master Bank",
		Username:    "admin",
		Password:    hashedMasterPassword, // You should hash the password before storing it in production
		Designation: "admin",
		Role:        "admin",
	}

	// Insert the master user into the database
	_, err = usersCollection.InsertOne(context.TODO(), masterUser)
	if err != nil {
		return err
	}

	log.Println("Master user initialized successfully.")
	return nil
}
