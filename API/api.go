package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

// GetUsers handles GET requests to fetch all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed. This is a GET method", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// CreateUser handles POST requests to create a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed. This is a POST method", http.StatusMethodNotAllowed)
		return
	}

	var newUser User
	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	users = append(users, newUser)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func main() {
	http.HandleFunc("/users", GetUsers)
	http.HandleFunc("/users/create", CreateUser)
	http.ListenAndServe(":8080", nil)
}
