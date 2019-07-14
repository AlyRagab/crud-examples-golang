package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var ctx, _ = context.WithTimeout(context.Background(), 60*time.Second)
var mongoClient *mongo.Client
var err error

func createUser(response http.ResponseWriter, request *http.Request) {
	// The Response Type
	response.Header().Add("content-type", "application/json")
	var user User
	json.NewDecoder(request.Body).Decode(&user)
	db := mongoClient.Database("userdb").Collection("user")
	result, err := db.InsertOne(ctx, db)
	check(err)
	json.NewEncoder(response).Encode(result)

}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	//Connection to MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err = mongo.Connect(ctx, clientOptions)
	check(err)
	fmt.Println("Database Connected Successfully !!")

	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/user", createUser).Methods("POST")
	err = http.ListenAndServe(":7070", myRouter)
}

