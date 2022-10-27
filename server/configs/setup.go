package configs

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var uri string = "mongodb+srv://Bases2:SiSale@cluster0.imibgld.mongodb.net/?retryWrites=true&w=majority"

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(
		options.Client().ApplyURI(uri),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

// Client instance
var DB *mongo.Client = ConnectDB()

// Getting database collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("metaOS").Collection(collectionName)
	return collection
}