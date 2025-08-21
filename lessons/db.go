package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv" // Importing the godotenv package to load environment variables
	_ "github.com/lib/pq"      // Importing the pq driver for PostgreSQL
)

var db *sql.DB

type User struct {
	ID   int
	Name string
	Age  int
}

func dbFunction() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}

	// Ping the database to verify connection
	if err := db.Ping(); err != nil {
		fmt.Println("Unable to reach the database:", err)
		return
	}
	fmt.Println("Successfully connected and pinged the database!")

	// defer db.Close() to close the database connection when the function exits
}

// createUser inserts a new user into the database and returns the user's ID.
func createUser(name string, age int) (int, error) {
	var id int
	query := `INSERT INTO users (name, age) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(query, name, age).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func readUsers() ([]User, error) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func readUserByID(id int) (*User, error) {
	var u User
	err := db.QueryRow("SELECT id, name, age FROM users WHERE id = $1", id).Scan(&u.ID, &u.Name, &u.Age)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func deleteUserByID(id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func updateUserByID(id int, name string, age int) error {
	result, err := db.Exec("UPDATE users SET name = $1, age = $2 WHERE id = $3", name, age, id)
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows == 0 {
		return fmt.Errorf("no user found with ID %d", id)
	}
	return nil
}
