package database

import (
	"context"
	"errors"

	"regexp"
	"strings"

	"github.com/suriya1776/hinata/usermangement-service/models"
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
	hasSpace := false

	for _, char := range password {
		if char >= 'A' && char <= 'Z' {
			hasUpperCase = true
		} else if char >= '0' && char <= '9' {
			hasNumber = true
		} else if strings.ContainsAny(string(char), "!@#$%^&*()-_=+[]{}|;:'\",.<>?/") {
			hasSpecialChar = true
		} else if char == ' ' {
			hasSpace = true
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
	if hasSpace {
		missingRequirements = append(missingRequirements, "spaces are not allowed")
	}

	if len(missingRequirements) > 0 {
		return errors.New("password requirements not met: " + strings.Join(missingRequirements, ", "))
	}

	return nil
}

// You need to implement these functions based on your preferred password hashing library
func HashPassword(password string) (string, error) {
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

// isValidBankName checks if the bank name is valid (not empty, <= 20 characters, no special characters)
func IsValidBankName(bankName string) bool {
	// Check if the bank name is not empty
	if bankName == "" {
		return false
	}

	// Check if the bank name length is within the allowed limit (20 characters)
	if len(bankName) > 20 {
		return false
	}

	// Check if the bank name contains only alphanumeric characters
	for _, char := range bankName {
		if (char < 'A' || char > 'Z') && (char < 'a' || char > 'z') && (char < '0' || char > '9') {
			return false
		}
	}
	return true
}

func containsSpecialCharacters(s string) bool {
	// Define a regular expression pattern for special characters
	// You can customize this pattern based on your definition of special characters
	// This example pattern allows any character that is not a letter or a number
	pattern := regexp.MustCompile(`[^a-zA-Z0-9]`)

	// Check if the string contains any special characters
	return pattern.MatchString(s)
}

// generatePasswordResetToken generates a unique token for password reset
// Generate a password reset token and send the reset email
