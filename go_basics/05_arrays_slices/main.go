package main

import "fmt"

func main() {
	// Declare Array and Assign Separately
	var fruitArr [2]string

	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	// Declare Array and Assign at same time
	bookArr := [2]string{"Midsummer Dream", "Hotspot"}

	// Slices
	citySlice := []string{"Singapore", "Thailand", "Taiwan"}

	fmt.Println(fruitArr, bookArr, citySlice)
	fmt.Println(len(citySlice))
	fmt.Println(citySlice[1:2])
}
