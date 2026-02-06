package data

import (
	"database/sql"
	"dbdemo/pkg/util"

	_ "modernc.org/sqlite"
)

func DBConnection() *sql.DB {
	db, err := sql.Open("sqlite", "file:companydb.db")
	util.ErrorHandling(err)
	return db
}

var CREATE_TABLE string = `CREATE TABLE IF NOT EXISTS employee(
	Id INTEGER PRIMARY KEY AUTOINCREMENT,
	name text,
	email text
	)`
func CreateTable(db *sql.DB,query string){

	_, err := db.Exec(query)
	util.ErrorHandling(err)

}
	
