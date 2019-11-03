package main

import "fmt"

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex 128

	// Using var
	var name string = "Brad"
	var age int = 32
	var salary int32 = 200
	var isCool = true
	const isHot = false // const cannot redefine unlike var
	// Shorthand
	nickname := "Brad"
	size := 1.5
	nick, email := "lim", "lim@gmail.com"

	fmt.Println(name, age, salary, isCool, isHot, nickname, size, nick, email)
	fmt.Printf("%T\n", size)
}
