package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Users endpoint")
    })

    http.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "Orders endpoint")
    })

    http.ListenAndServe(":8080", nil)
}
