package main

import "fmt"

func main() {
	// Define map
	emails := make(map[int]string)

	// Assign Key Value
	emails[1] = "bob@gmail.com"
	emails[2] = "sharon@gmail.com"
	emails[3] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails[2])

	// Delete from map
	delete(emails, 3)
	fmt.Println(emails)

	// Declare Map and add key value
	records := map[string]string{"Bob": "bob@gmail.com", "Sharon": "sharon@gmail.com"}
	fmt.Println(records)
}
