package repository

import (

	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"teams-secure-notes/internal/model"
)

var DB *mongo.Database

// InitializeDB initializes the DB connection
func InitializeDB(user, password, host, port, dbName string) {
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s/?authSource=admin&connect=direct",
		user, password, host, port, dbName)

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Failed to create a new MongoDB client: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %s", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %s", err)
	}

	DB = client.Database(dbName)
	log.Println("Successfully connected to MongoDB")
}

// CreateUser inserts a new user into the database
func CreateUser(user *model.User) error {
	collection := DB.Collection("users")
	log.Println(user)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Could not insert user: %v", err)
		return err
	}

	return nil
}

// GetUserByUsername finds a user by their username
func GetUserByUsername(username string) (*model.User, error) {
	collection := DB.Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user model.User
	err := collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil // No user found, return nil without an error
		}
		log.Printf("Could not find user: %v", err)
		return nil, err
	}
	return &user, nil
}
