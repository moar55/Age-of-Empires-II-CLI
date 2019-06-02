package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupDB() *mongo.Collection {
	dbURI := os.Getenv("DATABASE_URI")
	if dbURI == "" {
		log.Fatal("missing DB URI!")
	}
	// Set client options
	clientOptions := options.Client().ApplyURI(dbURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	if verbose {
		fmt.Println("Connected to MongoDB!")
	}
	//return the collection units of an age-of-empires db, this creates the db and collection if either aren't present
	return client.Database("age-of-empires").Collection("units")

}
