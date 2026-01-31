package main

import (
    "fmt"
    "net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Println("Request:", r.URL.Path)
        next.ServeHTTP(w, r)
    })
}

func hello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello with middleware")
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/hello", hello)

    http.ListenAndServe(":8080", loggingMiddleware(mux))
}
