package main

import "fmt"

func main() {
	var a = 10
	var b = 10
	var d = 5
	var e = 2
	var c = a + b - d * e
	fmt.Println(c)

	// augmented assignment
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)
	i += 5 // i + 5
	fmt.Println(i)

	// unary operator
	var j = 10
	j++
	fmt.Println(j)
	j++
	fmt.Println(j)
	
	j--
	fmt.Println(j)
	j--
	fmt.Println(j)
}