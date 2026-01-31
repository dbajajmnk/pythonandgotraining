// File: day14_middleware.go
// Topic: Middleware Concept using net/http
// Run: go run day14_middleware.go

package main

import (
	"context"
	"log"
	"net/http"
	"time"
)

type ctxKey string

const requestIDKey ctxKey = "request_id"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	handler := loggingMiddleware(requestIDMiddleware(authMiddleware(mux)))

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(requestIDKey)
	w.Write([]byte("Hello! request_id=" + id.(string)))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Println("Started", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Println("Completed in", time.Since(start))
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), requestIDKey, "req-12345")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
