package main

import "fmt"

func main() {
	// Define map
	emails := make(map[string]string)

	// Add key-value pairs to map
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	// Print map
	fmt.Println(emails)

	// Print map length
	fmt.Println(len(emails))

	// Print value of a key
	fmt.Println(emails["Bob"])

	// Delete from map
	delete(emails, "Bob")
	fmt.Println(emails)

	// Declare map and add key-value pairs
	emails2 := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com", "Mike": "mike@gmail.com"}
	fmt.Println(emails2)

	// Print map length
	fmt.Println(len(emails2))

	// Print value of a key
	fmt.Println(emails2["Bob"])

	// Delete from map
	delete(emails2, "Bob")
	fmt.Println(emails2)


}
