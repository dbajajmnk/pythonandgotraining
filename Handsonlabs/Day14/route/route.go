package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type car struct {
	NAME   string `json:"name"`
	Number string `json:"number"`
	Model  string `json:"model"`
}

var carsData = []car{car{NAME: "BMW", Number: "1", Model: "V1"},
	car{NAME: "BMW2", Number: "2", Model: "V2"}}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/cars", getCarsHandler)
	mux.HandleFunc("/cars/", carPathHandler)
	log.Println("Printing the logs of API")
	log.Fatal(http.ListenAndServe(":8080", loggingMiddleWare(mux)))

}

func healthHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Health is called")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
func getCarsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get Users is called")
	w.Header().Set("content-type", "json")
	json.NewEncoder(w).Encode(carsData)

}
func carPathHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Println("Path", path)
	pathAfterTrim := strings.Trim(path, "/")
	fmt.Println("Path After Trim", pathAfterTrim)
	parts := strings.Split(pathAfterTrim, "/")

	if len(parts) != 2 {
		http.NotFound(w, r)
		return
	}

	id := parts[1]
	fmt.Fprintf(w, "User ID: %s", id)

}

func loggingMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
		log.Printf("Request Path : %s", r.URL.Path)
		next.ServeHTTP(w, r)
	})

}
