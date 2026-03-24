package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", handler).Methods("GET")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users", createUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func getUsers(w http.ResponseWriter, r *http.Request) { // r -> pointer to request
	json.NewEncoder(w).Encode(users)
}

func createUser(w http.ResponseWriter, r *http.Request) { 
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser) // &newUser -> pointer to user
	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}

// Basic server
//
//	func main() {
//		http.HandleFunc("/", handler)
//		log.Fatal(http.ListenAndServe(":8080", nil))
//	}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World~!")
}
