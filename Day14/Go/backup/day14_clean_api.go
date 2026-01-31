package main

import (
    "encoding/json"
    "net/http"
)

type User struct {
    Name string `json:"name"`
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(User{Name: "Deepak"})
}

func main() {
    http.HandleFunc("/user", userHandler)
    http.ListenAndServe(":8080", nil)
}
