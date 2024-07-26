package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Lang struct {
	En string `bson:"en"`
	Th string `bson:"th"`
}

type Tier struct {
	ID   int  `bson:"id"`
	Name Lang `bson:"name"`
}

type User struct {
	ID       int    `bson:userid`
	Username string `bson:"username"`
	Password string `bson:"password"`
	Tier     *Tier  `bson:"tier"`
}

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

	testcol2 := client.Database("testdb").Collection("testcol2")

	// // Insert a document (new user)
	userAlice := User{
		Username: "useralice",
		Password: "passalice",
		Tier: &Tier{
			ID: 1,
			Name: Lang{
				En: "Gold",
				Th: "ทอง",
			},
		},
	}
	insertResult, err := testcol2.InsertOne(context.TODO(), userAlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult)

	// Find a user
	var result2 User
	err = collection.FindOne(context.TODO(), bson.D{{"username", userAlice.Username}}).Decode(&result2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result2)

	insertResult2, err := testcol2.InsertOne(context.TODO(), userAlice)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult2)

	// Find a user
	var result User
	err = collection.FindOne(context.TODO(), bson.D{{"username", userAlice.Username}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)

	// // Update a user
	// filter := bson.D{{"username", "exampleuser"}}
	// update := bson.D{
	// 	{"$set", bson.D{
	// 		{"password", "newexamplepass"},
	// 		{"tier", &Tier{
	// 			ID: 2,
	// 			Name: Lang{
	// 				En: "Platinum",
	// 				Th: "แพลตตินัม",
	// 			},
	// 		}},
	// 	}},
	// }
	// updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// // Delete a user
	// deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Deleted %v documents in the collection\n", deleteResult.DeletedCount)
}
