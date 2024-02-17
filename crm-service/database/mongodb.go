package database

import (
	"context"
	"errors"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/suriya1776/hinata/crm-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var (
	// client          *mongo.Client
	usersCollection *mongo.Collection
)

var (
	ErrBankExists   = errors.New("bank name already exists")
	ErrWeakPassword = errors.New("password does not meet strength requirements")
)

func init() {
	// Set your MongoDB Atlas connection string
	connectionURI := "mongodb+srv://system:manager@hinata.asdzn8q.mongodb.net"

	// Create a new MongoDB client
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// Select the database and collection
	usersCollection = client.Database("hinata").Collection("users")
}

// CreateUser creates a new user in the "users" collection
// CreateUser creates a new user in the "users" collection
func RegisterUser(c *gin.Context, user models.BankUser) error {
	// Check if the bank name already exists
	count, err := usersCollection.CountDocuments(context.TODO(), bson.M{"bankName": user.BankName})
	if err != nil {
		return err
	}
	if count > 0 {
		return ErrBankExists
	}

	// Perform password strength check
	err = isStrongPassword(user.Password)
	if err != nil {
		// Return the error without attempting to handle it here
		return err
	}

	// Hash the password before inserting
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	// Insert the user
	_, err = usersCollection.InsertOne(context.TODO(), user)
	return err
}

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
