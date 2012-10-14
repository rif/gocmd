package main

import (
    _ "github.com/bmizerany/pq"
    "database/sql"
    "log"
)

func main() {
    db, err := sql.Open("postgres", "user=rif dbname=gosqltest password=testus sslmode=disable")
    defer db.Close()
    if err != nil {
		log.Fatalf("failed to open the database: %v", err)
	}
	_, err = db.Exec("INSERT INTO cities VALUES ('Timisoara', '(23.0, 46.0)')")
	if err != nil {
		log.Fatalf("failed to execute insert statement: %v", err)
	} 		
	
    rows, err := db.Query("SELECT * FROM cities")
	if err != nil {
		log.Fatalf("failed to enumerate cities: %v", err)
	}
	for rows.Next() {
		var name, location string	
		rows.Scan(&name, &location)
		log.Print(name, location)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
