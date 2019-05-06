package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	people = append(people, Person{ID: "1", Firstname: "John", Lastname: "Doe", Address: &Address{City: "City X", State: "State X"}})
	people = append(people, Person{ID: "2", Firstname: "Koko", Lastname: "Doe", Address: &Address{City: "City Z", State: "State Y"}})
	people = append(people, Person{ID: "3", Firstname: "Francis", Lastname: "Sunday"})

	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}

// Person Description
type Person struct {
	ID        string   `json:"id,omitempty"`
	Firstname string   `json:"firstname,omitempty"`
	Lastname  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

// GetPeople list of people
func GetPeople(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(people)
	fmt.Println("w written")

}

// GetPerson Query for Person
func GetPerson(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			tempVar, _ := json.Marshal(item)
			fmt.Println(string(tempVar))
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// CreatePerson create person
func CreatePerson(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)

}

// DeletePerson delete person
func DeletePerson(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[(index+1):]...)
			break
		}
		json.NewEncoder(w).Encode(people)
	}
}
