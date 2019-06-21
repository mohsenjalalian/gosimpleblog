package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client


type Article struct {
	Title string `json:"Title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

type Articles []Article

func init() {
	client, _ = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	fmt.Println("This will get called on main initialization")
  }

func main()  {
	handleRequests()	
}

func allArticles(w http.ResponseWriter, req *http.Request) {
	articles := Articles {Article {Title: "test title", Desc: "test desc", Content: "test content"}}
	fmt.Println("articles")
	json.NewEncoder(w).Encode(articles)
}

func homePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "home")
}

func testPostArticles(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "post")
}


func CreatePersonEndpoint(	response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person Person
	json.NewDecoder(request.Body).Decode(&person)
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	result, _ := collection.InsertOne(context.TODO(), person)
	json.NewEncoder(response).Encode(result)
}



func GetPeopleEndpoint(response http.ResponseWriter, request *http.Request) { }
func GetPersonEndpoint(response http.ResponseWriter, request *http.Request) { }

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles).Methods("Get")
	myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	myRouter.HandleFunc("/person", CreatePersonEndpoint).Methods("POST")
	log.Fatal(http.ListenAndServe(":7777", myRouter))
}