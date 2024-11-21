package config

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DatabaseConfig struct {
	Client         *mongo.Client
	UserCollection *mongo.Collection
	Ctx            context.Context
}

func ConnectDatabase(config *AppConfig) *DatabaseConfig {
	ctx := context.TODO()

	clientOptions := options.Client().ApplyURI(config.DatabaseURI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("Error while connecting to MongoDB: ", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("Error while pinging MongoDB: ", err)
	}

	log.Println("Connected to MongoDB successfully")

	return &DatabaseConfig{
		Client:         client,
		UserCollection: client.Database("userdb").Collection("users"),
		Ctx:            ctx,
	}
}
