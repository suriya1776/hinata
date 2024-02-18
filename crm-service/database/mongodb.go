package database

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// client          *mongo.Client
	usersCollection, bankStateCollection *mongo.Collection
)

var (
	ErrBankExists   = errors.New("bank name already exists")
	ErrWeakPassword = errors.New("password does not meet strength requirements")
)

func init() {

	log.Println("starting db function")
	// Set your MongoDB Atlas connection string
	connectionURI := "mongodb+srv://system:manager@hinata.asdzn8q.mongodb.net"

	// Create a new MongoDB client
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Error pinging MongoDB:", err)
	}

	// Select the database and collection
	usersCollection = client.Database("hinata").Collection("users")
	bankStateCollection = client.Database("hinata").Collection("bank_state")

}
