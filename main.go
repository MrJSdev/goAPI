package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Student struct {
	ID         int    "json:id"
	FirstName  string "json:firstname"
	LastName   string "json:lastname"
	FatherName string "json:fathername"
}

var students []Student

func getStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	_ = json.NewDecoder(r.Body).Decode(&student)
	student.ID = rand.Intn(10000)
	students = append(students, student)
	json.NewEncoder(w).Encode(student)
}
func main() {
	name := "Shahnawaz Khan"
	fmt.Println(name)
	students = append(students, Student{ID: 2, FirstName: "Shahnawaz", LastName: "Khan", FatherName: "Shakeel Khan"})
	students = append(students, Student{ID: 3, FirstName: "Asif", LastName: "Ansari", FatherName: "Uknown Khan"})

	r := mux.NewRouter()
	r.HandleFunc("/api/students/", getStudents).Methods("GET")
	r.HandleFunc("/api/students/{id}", getStudent).Methods("GET")
	r.HandleFunc("/api/students/add", addStudent).Methods("POST")
	log.Fatal(http.ListenAndServe(":8001", r))
}
