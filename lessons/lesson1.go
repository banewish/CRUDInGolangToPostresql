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
		default:
			fmt.Println("Unknown choice.")
		}
	}
}
