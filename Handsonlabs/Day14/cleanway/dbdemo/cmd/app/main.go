package main

import (
	
	"dbdemo/internal/data"
	"fmt"
	"dbdemo/pkg/util"
	_ "modernc.org/sqlite"
)

func main() {
	fmt.Println("Data Base Example")
	db := data.DBConnection()

	//data.ErrorHandling(err)

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
