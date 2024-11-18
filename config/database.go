package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// DatabaseConfig holds the MongoDB client and collection
type DatabaseConfig struct {
	Client     *mongo.Client
	UserCollection *mongo.Collection
	Ctx        context.Context
}

// ConnectDatabase establishes a MongoDB connection
func ConnectDatabase(config *AppConfig) *DatabaseConfig {
	ctx := context.TODO()

	// Set up MongoDB client options
	clientOptions := options.Client().ApplyURI(config.DatabaseURI)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error while connecting to MongoDB: ", err)
	}

	// Verify the connection
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Error while pinging MongoDB: ", err)
	}

	log.Println("Connected to MongoDB successfully")

	// Return database configuration
	return &DatabaseConfig{
		Client:         client,
		UserCollection: client.Database("userdb").Collection("users"),
		Ctx:            ctx,
	}
}
