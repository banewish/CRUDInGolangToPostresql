package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dbFunction()
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Choose an action:")
		fmt.Println("1 - Create User")
		fmt.Println("2 - Read Users")
		fmt.Println("3 - Read User by ID")
		fmt.Println("4 - Delete User by ID")
		fmt.Println("0 - Exit")
		fmt.Print("Enter choice: ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter name: ")
			scanner.Scan()
			name := scanner.Text()

			fmt.Print("Enter age: ")
			scanner.Scan()
			ageStr := scanner.Text()
			age, err := strconv.Atoi(strings.TrimSpace(ageStr))
			if err != nil {
				fmt.Println("Invalid age")
				continue
			}

			id, err := createUser(name, age)
			if err != nil {
				fmt.Println("Error creating user:", err)
			} else {
				fmt.Printf("User created with ID: %d\n", id)
			}

		case "0":
			if db != nil {
				db.Close()
			}
			fmt.Println("Exiting.")
			return

		case "2":
			users, err := readUsers()
			if err != nil {
				fmt.Println("Error reading users:", err)
			} else {
				fmt.Println("Users:")
				for _, user := range users {
					fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
				}

			}
		case "3":
			fmt.Print("Enter user ID to read: ")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}
			user, err := readUserByID(id)
			if err != nil {
				fmt.Println("Error reading user:", err)
			} else {
				fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
			}

		case "4":
			fmt.Println("Enter user ID to delete:")
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(strings.TrimSpace(idStr))
			if err != nil {
				fmt.Println("Invalid ID")
				continue
			}
			err = deleteUserByID(id)
			if err != nil {
				fmt.Println("Error deleting user:", err)
			} else {
				fmt.Printf("User deleted with ID: %d\n", id)
			}

		default:
			fmt.Println("Unknown choice.")
		}
	}
}
