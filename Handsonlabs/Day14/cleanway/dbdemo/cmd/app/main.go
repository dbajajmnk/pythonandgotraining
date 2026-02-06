package main

import (
	
	"dbdemo/internal/data"
	"fmt"
	"dbdemo/pkg/util"
	"encoding/json"
	"log"
	"net/http"
	"database/sql"
	_ "modernc.org/sqlite"
)

type Server struct {
	db *sql.DB
}
func main() {
	fmt.Println("Data Base Example")
	db := data.DBConnection()

	//data.ErrorHandling(err)
	s := &Server{db: db}

	mux := http.NewServeMux()
	mux.HandleFunc("/health", s.health)

	addr := ":8080"
	log.Println("HTTP API running on http://localhost" + addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	CREATE_TABLE := `CREATE TABLE IF NOT EXISTS employee(
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	name text,
	email text
	)`
	data.CreateTable(db,CREATE_TABLE)
	
	fmt.Println("Table also Created")

	INSERT_EMPLOYEE_DATA := `INSERT INTO employee (name,email) VALUES (?,?)`
	_, err := db.Exec(INSERT_EMPLOYEE_DATA, "Deepak", "deepakbajaj79@gmail.com")

	util.ErrorHandling(err)

	fmt.Printf("Data is Inserted Successfully")
	READ_DATA := `SELECT Id,name,email from employee`

	rows, error := db.Query(READ_DATA)
	util.ErrorHandling(error)
	defer rows.Close()
	fmt.Printf("Rows Fetched")

	for rows.Next() {
		var id int
		var name, email string
		rows.Scan(&id, &name, &email)
		fmt.Printf("Id :%d,Name : %s, Email:%s", id, name, email)
	}

}
func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, map[string]any{"error": "method not allowed"})
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"status": "ok"})
}

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
