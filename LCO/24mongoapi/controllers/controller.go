package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"
)

const connectionString = "mongodb+srv://sukanta:suka207@cluster0.yw4mca7.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const collectionName = "watchlist"

// MOST IMPORTANT
var collection *mongo.Collection

// connect with mongoDB

// runs at the start and run only once
func init() {
	// client connection
	clientOption := options.Client().ApplyURI(connectionString)
	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(collectionName)

	// collection instance
	fmt.Println("Collection instance ready")
}
