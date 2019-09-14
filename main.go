package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Student struct {
	ID         int    `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName  string `json:"firstname,omitempty" bson:"firstname,omitempty"`
	LastName   string `json:"lastname,omitempty" bson:"lastname,omitempty"`
	FatherName string `json:"fathername,omitempty" bson:"fathername,omitempty"`
}

var students []Student

// Connect to MongoDB
var client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
var collection = client.Database("sahara").Collection("students")

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	students, err = collection.Find()
	json.NewEncoder(w).Encode(students)
}
func getStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, item := range students {
		if strconv.Itoa(item.ID) == param["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}
}
func addStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var student Student
	json.NewDecoder(r.Body).Decode(&student)
	result, err := collection.InsertOne(context.TODO(), student)
	// students = append(students, student)
	fmt.Println(err)
	json.NewEncoder(w).Encode(result)
}
func main() {
	name := "Shahnawaz Khan"
	fmt.Println(name)
	// students = append(students, Student{ID: 2, FirstName: "Shahnawaz", LastName: "Khan", FatherName: "Shakeel Khan"})
	// students = append(students, Student{ID: 3, FirstName: "Asif", LastName: "Ansari", FatherName: "Uknown Khan"})
	r := mux.NewRouter()
	r.HandleFunc("/api/students/", getStudents).Methods("GET")
	r.HandleFunc("/api/students/{id}", getStudent).Methods("GET")
	r.HandleFunc("/api/students/add", addStudent).Methods("POST")

	log.Fatal(http.ListenAndServe(":8001", r))

}
