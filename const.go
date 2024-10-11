package main

import "fmt"

func main() {
	const nama string = "Nabildzr"
	fmt.Println("Hello, ", nama)

	const (
		firstName = "Nabil"
		lastName = "Dzikrika"
	)

	fmt.Println("Hello, ", firstName, lastName)

	// Error: cannot assign to nama
	// nama = "Argow"
	// fmt.Println("Hello, ", nama)
}