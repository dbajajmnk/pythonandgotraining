// File: day14_routing_basics.go
// Topic: Routing Basics using net/http
// Run: go run day14_routing_basics.go

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/users/", userByIDHandler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Users list"))
}

func userByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Manual path param parsing: /users/{id}
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) != 2 {
		http.NotFound(w, r)
		return
	}
	id := parts[1]
	fmt.Fprintf(w, "User ID: %s", id)
}
