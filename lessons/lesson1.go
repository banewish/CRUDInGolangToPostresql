package main

func main() {

	dbFunction()

	createUser("alex", 30)
	if db != nil {
		db.Close() // Close DB after all operations
	}
}
