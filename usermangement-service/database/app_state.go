// database/bank_state.go

package database

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/suriya1776/hinata/usermangement-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrBankNotFound = errors.New("bank not found")
	rolesCollection *mongo.Collection
)

// UpdateBankNameRequest represents the request structure for updating the bank name
type UpdateBankNameRequest struct {
	NewBankName string `json:"newBankName" binding:"required"`
}

// UpdateBankName updates the bank name in the bank_state collection
func UpdateBankName(newBankName string) error {
	filter := bson.M{} // You can add more specific conditions here if needed

	update := bson.M{"$set": bson.M{"bankName": newBankName}}

	result, err := bankStateCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Println("Error updating bank name:", err)
		return err
	}

	if result.ModifiedCount == 0 {
		return ErrBankNotFound
	}

	return nil
}

// InsertRole inserts a new role into the "roles" collection
func InsertRole(role models.Role) error {
	// Check if the role name is empty
	if len(role.Name) == 0 {
		return errors.New("role name cannot be empty")
	}

	// Check if the role name is greater than 10 characters
	if len(role.Name) > 10 {
		return errors.New("role name cannot be greater than 10 characters")
	}

	// Check if the role name contains special characters
	if containsSpecialCharacters(role.Name) {
		return errors.New("role name cannot contain special characters")
	}

	// Check if the role already exists
	count, err := rolesCollection.CountDocuments(context.TODO(), bson.M{"name": role.Name})
	if err != nil {
		return err
	}
	if count > 0 {
		log.Println("Role already exists. Skipping insertion.")
		return nil
	}

	// Insert the role into the "roles" collection
	_, err = rolesCollection.InsertOne(context.TODO(), role)
	return err
}

// IsValidRole checks if the provided role name is valid
func IsValidRole(roleName string) bool {
	// Your logic for role name validation (e.g., no special characters, length checks)
	return !containsSpecialCharacters(roleName) && len(roleName) > 0 && len(roleName) <= 20
}

// GrantRoles grants the specified roles to a user
func GrantRoles(username string, roles []string) error {
	// Check if the target user exists
	user, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	// Validate each role before granting
	for _, role := range roles {
		role = strings.ToUpper(role) // Convert role name to uppercase
		if !IsValidRole(role) {
			return errors.New("Invalid role name: " + role)
		}
	}

	// Add roles to the user document
	update := bson.M{"$addToSet": bson.M{"roles": bson.M{"$each": roles}}}
	_, err = usersCollection.UpdateOne(context.TODO(), bson.M{"username": username}, update)
	if err != nil {
		log.Println("Failed to update user roles:", err)
		return errors.New("failed to update user roles")
	}

	return nil
}
