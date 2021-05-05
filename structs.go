package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	city      string
	gender    string
	age       int
}

func main() {
	person1 := person{firstname: "Example", lastname: "Example", city: "Example", gender: "M", age: 20}

	fmt.Println(person1)
}
