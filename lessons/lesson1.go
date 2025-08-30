package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parsePositiveInt(s string) (int, error) {
	n, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil || n <= 0 {
		return 0, fmt.Errorf("please enter a positive number")
	}
	return n, nil
}

func isValidName(name string) bool {
	return strings.TrimSpace(name) != ""
}

func isValidCountryCode(code string) bool {
	return strings.TrimSpace(code) != ""
}

func main() {
	dbFunction()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Choose an action:")
		fmt.Println("1 - Create User")
		fmt.Println("2 - Read Users")
		fmt.Println("3 - Read User by ID")
		fmt.Println("4 - Delete User by ID")
		fmt.Println("5 - Update User by ID")
		fmt.Println("6 - Create Country")
		fmt.Println("7 - Create UserInfo")
		fmt.Println("8 - Update UserInfo")
		fmt.Println("0 - Exit")
		fmt.Print("Enter choice: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter username: ")
			scanner.Scan()
			name := scanner.Text()
			if !isValidName(name) {
				log.Println("Username cannot be empty")
				continue
			}

			fmt.Print("Enter country code: ")
			scanner.Scan()
			countryCode := scanner.Text()
			if !isValidCountryCode(countryCode) {
				log.Println("Country code cannot be empty")
				continue
			}

			id, err := createUser(name, countryCode)
			if err != nil {
				log.Println("Error creating user:", err)
			} else {
				fmt.Printf("User created with ID: %d\n", id)
			}

		case "2":
			users, err := readUsers()
			if err != nil {
				log.Println("Error reading users:", err)
			} else {
				fmt.Println("Users:")
				for _, user := range users {
					fmt.Printf("ID: %d, Username: %s, CountryCode: %s\n", user.ID, user.Username, user.CountryCode)
				}
			}

		case "3":
			fmt.Print("Enter user ID to read: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := parsePositiveInt(idStr)
			if err != nil {
				log.Println("Invalid ID")
				continue
			}

			user, countryName, err := readUserByID(id)
			if err != nil {
				log.Println("Error reading user:", err)
			} else {
				fmt.Printf("ID: %d, Username: %s, CountryCode: %s, CountryName: %s\n", user.ID, user.Username, user.CountryCode, countryName)
			}

		case "4":
			fmt.Print("Enter user ID to delete: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := parsePositiveInt(idStr)
			if err != nil {
				log.Println("Invalid ID")
				continue
			}
			err = deleteUserByID(id)
			if err != nil {
				log.Println("Error deleting user:", err)
			} else {
				fmt.Printf("User deleted with ID: %d\n", id)
			}

		case "5":
			fmt.Print("Enter user ID to update: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := parsePositiveInt(idStr)
			if err != nil {
				log.Println("Invalid ID")
				continue
			}

			fmt.Print("Enter new username: ")
			scanner.Scan()
			newName := scanner.Text()
			if !isValidName(newName) {
				log.Println("Username cannot be empty")
				continue
			}

			fmt.Print("Enter new country code: ")
			scanner.Scan()
			newCountryCode := scanner.Text()
			if !isValidCountryCode(newCountryCode) {
				log.Println("Country code cannot be empty")
				continue
			}

			err = updateUserByID(id, newName, newCountryCode)
			if err != nil {
				log.Println("Error updating user:", err)
			} else {
				fmt.Printf("User with ID %d updated successfully.\n", id)
			}

		case "6":
			fmt.Print("Enter country code: ")
			scanner.Scan()
			code := scanner.Text()
			if !isValidCountryCode(code) {
				log.Println("Country code cannot be empty")
				continue
			}
			fmt.Print("Enter country name: ")
			scanner.Scan()
			name := scanner.Text()
			if !isValidName(name) {
				log.Println("Country name cannot be empty")
				continue
			}
			err := createCountry(code, name)
			if err != nil {
				log.Println("Error creating country:", err)
			} else {
				fmt.Println("Country created successfully.")
			}

		case "7":
			fmt.Print("Enter user ID for userInfo: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := parsePositiveInt(idStr)
			if err != nil {
				log.Println("Invalid ID")
				continue
			}
			fmt.Print("Enter user phone: ")
			scanner.Scan()
			phone := scanner.Text()
			if !isValidName(phone) {
				log.Println("Phone cannot be empty")
				continue
			}
			err = createUserInfo(id, phone)
			if err != nil {
				log.Println("Error creating userInfo:", err)
			} else {
				fmt.Println("UserInfo created successfully.")
			}

		case "8":
			fmt.Print("Enter user ID for userInfo update: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := parsePositiveInt(idStr)
			if err != nil {
				log.Println("Invalid ID")
				continue
			}
			fmt.Print("Enter new user phone: ")
			scanner.Scan()
			phone := scanner.Text()
			if !isValidName(phone) {
				log.Println("Phone cannot be empty")
				continue
			}
			err = updateUserInfo(id, phone)
			if err != nil {
				log.Println("Error updating userInfo:", err)
			} else {
				fmt.Println("UserInfo updated successfully.")
			}

		case "0":
			if db != nil {
				db.Close()
			}
			fmt.Println("Exiting.")
			return

		default:
			log.Println("Unknown choice.")
		}
	}
}
