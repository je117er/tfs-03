package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/tfs")
	if err != nil {
		// do something
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM lecs LIMIT 5")
	// buffer

	rows.Next() {

	}
	var a, b, c int
	rows.Scan(&a, &b, &c)
	defer rows.Close()

}
