package main

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

func main() {
	fmt.Println("Data Base Example")
	db, err := sql.Open("sqlite", "file:companydb.db")

	if err != nil {
		fmt.Println("Sorry! Unable to Connect with Database", err)
	}
	defer db.Close()
	CREATE_TABLE := `CREATE TABLE IF NOT EXISTS employee(
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	name text,
	email text
	)`
	_, err = db.Exec(CREATE_TABLE)
	if err != nil {
		fmt.Print("Table Creaton Failed", err)
	}
	fmt.Println("Table also Created")

	INSERT_EMPLOYEE_DATA := `INSERT INTO employee (name,email) VALUES (?,?)`
	_, err = db.Exec(INSERT_EMPLOYEE_DATA, "Deepak", "deepakbajaj79@gmail.com")

	if err != nil {
		fmt.Printf("Data Insert Failed")
	}

	fmt.Printf("Data is Inserted Successfully")
	READ_DATA := `SELECT Id,name,email from employee`

	rows, error := db.Query(READ_DATA)
	if err != nil {
		fmt.Print("Unable to Fetch Data", error)
	}
	defer rows.Close()
	fmt.Printf("Rows Fetched")

	for rows.Next(){
		var id int
		var name,email string
		rows.Scan(&id,&name,&email)
		fmt.Printf("Id :%d,Name : %s, Email:%s",id,name,email)
	}

}
