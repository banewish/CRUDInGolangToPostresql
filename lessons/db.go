package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Importing the pq driver for PostgreSQL
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "postgres"
)

func dbFunction() {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	defer db.Close()

	// Ping the database to verify connection
	if err := db.Ping(); err != nil {
		fmt.Println("Unable to reach the database:", err)
		return
	}
	fmt.Println("Successfully connected and pinged the database!")
}
