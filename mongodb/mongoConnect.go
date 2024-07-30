package mongodb

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserLogin struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username string             `json:"username,omitempty" bson:"username,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omitempty"`
	Tier     *Tier              `json:"tier,omitempty" bson:"tier,omitempty"`
}

type Lang struct {
	En string `json:"en,omitempty" bson:"en,omitempty"`
	Th string `json:"th,omitempty" bson:"th,omitempty"`
}

type Tier struct {
	Id   int  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name Lang `json:"name,omitempty" bson:"name,omitempty"`
}

var client *mongo.Client

func SetClient(mongoClient *mongo.Client) {
	client = mongoClient
}

func CreateUserLogin(c *fiber.Ctx) error {
	var user UserLogin
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	collection := client.Database("testdb2").Collection("login")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(result)
}

func GetAllUserLogin(c *fiber.Ctx) error {
	var users []UserLogin
	collection := client.Database("testdb2").Collection("login")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user UserLogin
		if err := cursor.Decode(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(users)
}

func GetUserLoginByUsername(c *fiber.Ctx) error {
	name := c.Params("username")
	var user UserLogin
	collection := client.Database("testdb2").Collection("login")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, bson.M{"username": name}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(user)
}

func GetUserLoginById(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID format")
	}

	var user UserLogin
	collection := client.Database("testdb2").Collection("login")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(user)
}

func UpdateUserLoginById(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID format")
	}

	var user UserLogin
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	collection := client.Database("testdb2").Collection("login")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}
	update := bson.M{"$set": user}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(result)
}

func DeleteUserLoginById(c *fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID format")
	}

	collection := client.Database("testdb2").Collection("login")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.JSON(result)
}

func ConnectToMongo() {
	fmt.Println("Starting connect to mongo database..")
	// define timeout
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://mongouser:mongopass@localhost:27017")
	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting to MongoDB:", err)
		return
	}
}

