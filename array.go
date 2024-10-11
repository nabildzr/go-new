package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Nabil"
	names[1] = "Dzikrika"
	names[2] = "Ramdani"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	
	var values = [...]int{
		90,
		80,
		95,
	
	}

	

	fmt.Println(values)
}