package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Contact struct {
	ID         int `json:"id"`
	Username   string `json:"username"`
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
	FullName   string `json:"fullName"`
	Phone      []struct {
		TypeID      int `json:"typeID"`
		CountryCode int `json:"countryCode"`
		Operator    int `json:"operator"`
		Number      int `json:"number"`
	} `json:"phone"`
	Email     []string  `json:"email"`
	Birthdate time.Time `json:"birthdate"`
}

type Group struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Contacts    []int  `json:"contacts"`
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	contact := Contact{}
	
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(contact)
	case "POST", "PUT":
		json.NewEncoder(w).Encode(contact)
	case "DELETE":
		json.NewEncoder(w).Encode(contact)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func groupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	group := Group{}
	
	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(group)
	case "POST", "PUT":
		json.NewEncoder(w).Encode(group)
	case "DELETE":
		json.NewEncoder(w).Encode(group)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/api/v1/contact", contactHandler)
	http.HandleFunc("/api/v1/group", groupHandler)

	fmt.Println("Server starting on port 6080...")
	http.ListenAndServe(":6080", nil)
}