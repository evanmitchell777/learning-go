package main

import "fmt"

func main() {
	age := 20
	if age > 60 {
		fmt.Print("Getting older")

	} else if age > 30 {
		fmt.Print("Getting wiser")
	} else if age > 20 {
		fmt.Print("Adulthood")
	} else if age > 10 {
		fmt.Print("young blood")
	} else {
		fmt.Print("Booting up")
	}
}
