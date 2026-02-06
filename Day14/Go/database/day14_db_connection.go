
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

func main() {
	fmt.Println("=== Day 14: SQLite DB Connection (Pure Go) ===")

	db, err := sql.Open("sqlite", "file:day14.db")
	if err != nil {
		log.Fatal("DB open error:", err)
	}
	defer db.Close()

	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal("Table creation failed:", err)
	}

	_, err = db.Exec(`INSERT INTO users(name, email) VALUES (?, ?)`, "Deepak", "deepak@example.com")
	if err != nil {
		log.Fatal("Insert failed:", err)
	}

	rows, err := db.Query(`SELECT id, name, email FROM users`)
	if err != nil {
		log.Fatal("Select failed:", err)
	}
	defer rows.Close()

	fmt.Println("Users:")
	for rows.Next() {
		var id int
		var name, email string
		rows.Scan(&id, &name, &email)
		fmt.Printf("ID=%d Name=%s Email=%s\n", id, name, email)
	}

	fmt.Println("=== DONE ===")
}
