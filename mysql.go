package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "cgrates:testus@tcp(192.168.0.17:3306)/cgrates?charset=utf8")
	defer db.Close()
	if err != nil {
		log.Fatalf("failed to open the database: %v", err)
	}
	_, err = db.Exec("INSERT INTO emp VALUES ('0000000001', 'Radu', 'Fericean')")
	if err != nil {
		log.Fatalf("failed to execute insert statement: %v", err)
	}

	rows, err := db.Query("SELECT * FROM emp")
	if err != nil {
		log.Fatalf("failed to enumerate employees: %v", err)
	}
	for rows.Next() {
		var id, firstname, lastname string
		rows.Scan(&id, &firstname, &lastname)
		log.Printf("%s %s %s", id, firstname, lastname)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
