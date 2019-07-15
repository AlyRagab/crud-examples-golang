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
	"gopkg.in/mgo.v2/bson"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var ctx, _ = context.WithTimeout(context.Background(), 20*time.Second)
var mongoClient *mongo.Client
var err error
var person Person
var people []Person

// Handling the Error function
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Create Users in the people Collection
func createUsers(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")

	json.NewDecoder(r.Body).Decode(&person)
	// Create the Database
	db := mongoClient.Database("usersdb").Collection("people")
	resultdb, err := db.InsertOne(ctx, db)
	check(err)
	// Encode the Json data
	json.NewEncoder(w).Encode(resultdb)

}

// Delete The Users in the people Collection
func deleteUsers(w http.ResponseWriter, r *http.Request) {

}

// Querying and Getting The users
func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")
	db := mongoClient.Database("usersdb").Collection("people")
	get, err := db.Find(ctx, bson.M{})
	check(err)
	defer get.Close(ctx)

	for get.Next(ctx) {
		var person Person
		get.Decode(&person)
		people = append(people, person)

	}
	json.NewEncoder(w).Encode(people)
}

// Insert Users in the people Collection
func insertUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

}

func main() {

	// Database Connection
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err = mongo.Connect(ctx, clientOptions)
	check(err)
	fmt.Println("Database Connected Successfully !!")

	//Setting up the Router
	myRouter := mux.NewRouter()
	myRouter.HandleFunc("/user", createUsers).Methods("POST")
	myRouter.HandleFunc("/user", deleteUsers).Methods("DELETE")
	myRouter.HandleFunc("/user", getUsers).Methods("GET")
	myRouter.HandleFunc("/user", insertUsers).Methods("PUT")
	err = http.ListenAndServe(":7070", myRouter)
}
