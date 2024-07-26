package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
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

type UserLogin struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username  string             `json:"username,omitempty" bson:"username,omitempty"`
	Password  string             `json:"password,omitempty" bson:"password,omitempty"`
	firstname string
	lastname  string
	Tier      *Tier `json:"tier,omitempty" bson:"tier,omitempty"`
}

type UserLogin struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Username   string             `json:"username,omitempty" bson:"username,omitempty"`
	Password   string             `json:"password,omitempty" bson:"password,omitempty"`
	firstname  string
	lastname   string
	middlename string
	Tier       *Tier `json:"tier,omitempty" bson:"tier,omitempty"`
}

var client *mongo.Client

func CreateUserLoginEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var user UserLogin
	json.NewDecoder(request.Body).Decode(&user)
	collection := client.Database("testdb2").Collection("login")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(response).Encode(result)
}

func GetAllUserLoginEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var users []UserLogin
	collection := client.Database("testdb2").Collection("login")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	curser, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
	}
	defer curser.Close(ctx)
	for curser.Next(ctx) {
		var user UserLogin
		curser.Decode(&user)
		users = append(users, user)
	}
	if err := curser.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(users)
}

func GetUserLoginByUsername(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	name := params["username"]
	var user UserLogin
	collection := client.Database("testdb2").Collection("login")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, bson.M{"username": name}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
	}
	json.NewEncoder(response).Encode(user)
}

func GetUserLoginById(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(`{ "message": "Invalid ID format"}`))
		return
	}

	var user UserLogin
	collection := client.Database("testdb2").Collection("login")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = collection.FindOne(ctx, UserLogin{ID: id}).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
	}
	json.NewEncoder(response).Encode(user)
}

func UpdateUserLoginById(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var user UserLogin
	json.NewDecoder(request.Body).Decode(&user)
	collection := client.Database("testdb2").Collection("login")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"_id": id}
	update := bson.M{
		"$set": user,
	}
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func DeleteUserLoginById(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	collection := client.Database("testdb2").Collection("login")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"_id": id}
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(result)
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
	router := mux.NewRouter()
	router.HandleFunc("/create-user", CreateUserLoginEndpoint).Methods("POST")
	router.HandleFunc("/user", GetAllUserLoginEndpoint).Methods("GET")
	router.HandleFunc("/user/{username}", GetUserLoginByUsername).Methods("GET")
	router.HandleFunc("/user-id/{id}", GetUserLoginById).Methods("GET")
	router.HandleFunc("/update-user/{id}", UpdateUserLoginById).Methods("PUT")
	router.HandleFunc("/delete-user/{id}", DeleteUserLoginById).Methods("DELETE")
	http.ListenAndServe(":8081", router)
}
