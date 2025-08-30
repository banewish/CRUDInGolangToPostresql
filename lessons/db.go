package main

import (
	"database/sql"
	"fmt"
	"os"

	"encoding/json"
	"time"

	"github.com/joho/godotenv" // Importing the godotenv package to load environment variables
	_ "github.com/lib/pq"      // Importing the pq driver for PostgreSQL
)

var db *sql.DB

type User struct {
	ID          int
	Username    string
	CountryCode string
}

type Country struct {
	CountryCode string
	CountryName string
}

type UserInfo struct {
	ID        int
	UserPhone string
	Metadata  map[string]interface{}
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
func createUser(username, countryCode string) (int, error) {
	var id int
	query := `INSERT INTO users (username, countryCode) VALUES ($1, $2) RETURNING id`
	err := db.QueryRow(query, username, countryCode).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func readUsers() ([]User, error) {
	rows, err := db.Query("SELECT id, username, countryCode FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.CountryCode); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func readUserByID(userID int) (*User, string, error) {
	query := `
        SELECT u.id, u.username, u.countryCode, c.countryName
        FROM users u
        JOIN countries c ON u.countryCode = c.countryCode
        WHERE u.id = $1
    `
	row := db.QueryRow(query, userID)
	var user User
	var countryName string
	err := row.Scan(&user.ID, &user.Username, &user.CountryCode, &countryName)
	if err != nil {
		return nil, "", err
	}
	return &user, countryName, nil
}

func deleteUserByID(id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func updateUserByID(id int, username string, countryCode string) error {
	result, err := db.Exec("UPDATE users SET username = $1, countryCode = $2 WHERE id = $3", username, countryCode, id)
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

func createUserInfo(id int, phone string) error {
	metadata := map[string]interface{}{
		"createdAt": time.Now().Format(time.RFC3339),
		"updatedAt": time.Now().Format(time.RFC3339),
	}
	metaBytes, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	query := `INSERT INTO userInfo (id, userPhone, metadata) VALUES ($1, $2, $3)`
	_, err = db.Exec(query, id, phone, string(metaBytes))
	return err
}

func updateUserInfo(id int, phone string) error {
	metadata := map[string]interface{}{
		"updatedAt": time.Now().Format(time.RFC3339),
	}
	metaBytes, err := json.Marshal(metadata)
	if err != nil {
		return err
	}

	query := `UPDATE userInfo SET userPhone = $1, metadata = metadata || $2::jsonb WHERE id = $3`
	_, err = db.Exec(query, phone, string(metaBytes), id)
	return err
}

func createCountry(code, name string) error {
	_, err := db.Exec("INSERT INTO countries (countryCode, countryName) VALUES ($1, $2)", code, name)
	return err
}
