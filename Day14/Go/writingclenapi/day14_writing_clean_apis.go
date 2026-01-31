// File: day14_writing_clean_apis.go
// Topic: Writing Clean APIs in Go
// Run: go run day14_writing_clean_apis.go

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var users = []User{
	{ID: 1, Name: "Alice"},
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", usersHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		writeJSON(w, http.StatusOK, users)

	case http.MethodPost:
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			writeError(w, http.StatusBadRequest, "INVALID_JSON", "Invalid JSON payload")
			return
		}
		if u.Name == "" {
			writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "name is required")
			return
		}
		u.ID = len(users) + 1
		users = append(users, u)
		writeJSON(w, http.StatusCreated, u)

	default:
		writeError(w, http.StatusMethodNotAllowed, "METHOD_NOT_ALLOWED", "Method not allowed")
	}
}

func writeJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, ErrorResponse{
		Code:    code,
		Message: message,
	})
}
