// database/initialize.go
package database

import (
	"context"
	"log"

	"github.com/suriya1776/hinata/crm-service/models"
	"go.mongodb.org/mongo-driver/bson"
)

// InitializeDB inserts a master user into the database
func InitializeMasterUser() error {
	// Check if the master user already exists
	log.Println("Entering intilize DB")

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
		Username:    "admin",
		Password:    hashedMasterPassword, // You should hash the password before storing it in production
		Designation: "admin",
		Roles:       []string{"admin"},
	}

	// Insert the master user into the "users" collection
	_, err = usersCollection.InsertOne(context.TODO(), masterUser)
	if err != nil {
		return err
	}

	log.Println("Master user and bank state initialized successfully.")
	return nil
}

func InitializeBankState() error {
	// Check if the bank state already exists
	count, err := bankStateCollection.CountDocuments(context.TODO(), bson.M{"bankName": "Demo bank"})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Bank state already exists. Skipping initialization.")
		return nil
	}

	// Create an entry in the "bank_state" collection
	bankState := models.BankState{
		BankName: "Demo bank", // Set your desired bank name here
		// Add other bank-related fields as needed
	}

	// Insert the bank state into the "bank_state" collection
	_, err = bankStateCollection.InsertOne(context.TODO(), bankState)
	if err != nil {
		return err
	}

	log.Println("Bank state initialized successfully.")
	return nil
}
