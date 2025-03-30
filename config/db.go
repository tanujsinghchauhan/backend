package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database // Global variable for the database

// GetCollection returns a reference to a MongoDB collection
func GetCollection(collectionName string) *mongo.Collection {
	if DB == nil {
		log.Println("⚠️ Database not connected. Returning nil.")
		return nil
	}
	return DB.Collection(collectionName)
}

func ConnectDB() error { // Change return type to error
	// Create a new MongoDB client
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Println("❌ Database connection error:", err)
		return err
	}

	// Set timeout for connection verification
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Verify the connection with Ping
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Println("❌ Database not connected:", err)
		return err
	}

	// Set the database name
	DB = client.Database("event_management")
	fmt.Println("✅ Successfully connected to MongoDB locally!")
	return nil // Return nil if successful
}
