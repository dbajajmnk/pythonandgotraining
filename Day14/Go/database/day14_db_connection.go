package main

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/mattn/go-sqlite3"
)

func main() {
    db, err := sql.Open("sqlite3", "app.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }

    db.Exec("CREATE TABLE IF NOT EXISTS users(id INTEGER, name TEXT)")
    db.Exec("INSERT INTO users(id, name) VALUES(1, 'Alice')")

    rows, _ := db.Query("SELECT id, name FROM users")
    defer rows.Close()

    for rows.Next() {
        var id int
        var name string
        rows.Scan(&id, &name)
        fmt.Println(id, name)
    }
}
