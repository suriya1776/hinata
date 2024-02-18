package database

import (
	"context"
	"errors"

	"strings"

	"github.com/suriya1776/hinata/crm-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

// isStrongPassword checks if the password meets the strength criteria
// isStrongPassword checks if the password meets the strength criteria
func isStrongPassword(password string) error {
	// Require at least 8 characters, 1 uppercase, 1 number, and 1 special character
	minLength := 8
	hasUpperCase := false
	hasNumber := false
	hasSpecialChar := false

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpperCase = true
		} else if char >= '0' && char <= '9' {
			hasNumber = true
		} else if strings.ContainsAny(string(char), "!@#$%^&*()-_=+[]{}|;:'\",.<>?/") {
			hasSpecialChar = true
		}
	}

	if len(password) < minLength {
		return errors.New("password is too short, must be at least 8 characters")
	}

	var missingRequirements []string
	if !hasUpperCase {
		missingRequirements = append(missingRequirements, "at least 1 uppercase letter")
	}
	if !hasNumber {
		missingRequirements = append(missingRequirements, "at least 1 number")
	}
	if !hasSpecialChar {
		missingRequirements = append(missingRequirements, "at least 1 special character")
	}

	if len(missingRequirements) > 0 {
		return errors.New("password requirements not met: " + strings.Join(missingRequirements, ", "))
	}

	return nil
}

// You need to implement these functions based on your preferred password hashing library
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func GetUserByUsername(username string) (*models.BankUser, error) {
	var user models.BankUser

	// Create a filter to find the user by username
	filter := bson.M{"username": username}

	// Find the user in the database
	err := usersCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// No user found with the provided username
			return nil, nil
		}
		// An error occurred while fetching the user
		return nil, err
	}

	return &user, nil
}

// isValidUsername checks if the username is valid (only text and numbers allowed, no special characters)
func IsValidUsername(username string) bool {
	for _, char := range username {
		if !((char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9')) {
			return false
		}
	}
	return true
}

// isUsernameUnique checks if the username is unique in the database
func IsUsernameUnique(username string) error {
	user, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	if user != nil {
		return errors.New("username already exists")
	}
	return nil
}
