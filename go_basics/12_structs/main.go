package main

import (
	"fmt"
	"strconv"
)

// Person is a representation of a person
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting Method (value receiver) - used if not changing any data
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " " + p.city + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday Method (pointer receiver) - used if changing data
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{firstName: "Alan", lastName: "Smith", city: "TW", gender: "m", age: 20}
	// Alternativce Init
	person2 := Person{"Samantha", "Tan", "UK", "f", 25}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person1.firstName)

	person1.age = 30
	fmt.Println(person1)

	person1.hasBirthday()
	person2.getMarried("William")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())
}
