// File: day14_rest_net_http.go
// Topic: REST APIs using net/http
// Run: go run day14_rest_net_http.go

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone`
}

var users = []User{
	{ID: 1, Name: "Alice", Phone:"838388383883"},
	{ID: 2, Name: "Bob", Phone :"234737737737"},
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/print",printUserRequest)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", loggingMiddleware(mux)))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func printUserRequest(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
    var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(u)
	

}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(users)

	case http.MethodPost:
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "invalid json", http.StatusBadRequest)
			return
		}
		u.ID = len(users) + 1
		users = append(users, u)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(u)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}