package main

import "fmt"

// if no return type, don't put anything into return
func greeting(name string) string {
	return "Hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greeting("Tim"))
	fmt.Println(getSum(10, 12))
}
