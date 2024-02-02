package repository

import (
	"context"
	"fmt"
	"log"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

