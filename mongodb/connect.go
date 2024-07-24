package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectMongo() {
	clientOptions := options.Client().ApplyURI("mongodb://mongouser:mongopass@localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// Get a handle for the collection
	collection := client.Database("testdb").Collection("testcol")

	// Insert a document
	doc := bson.D{{"name", "Alice"}, {"age", 25}}
	insertResult, err := collection.InsertOne(context.TODO(), doc)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a document: ", insertResult.InsertedID)

	// Find a document
	var result bson.D
	err = collection.FindOne(context.TODO(), bson.D{{"name", "Alice"}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Found a document: ", result)
}
