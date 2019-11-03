package main

import "fmt"

func main() {
	a := 5
	b := &a

	// everything in go is passed by value

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// Use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// Change val with pointer
	*b = 10
	fmt.Println(a)
}
