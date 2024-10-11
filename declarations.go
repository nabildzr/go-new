package main

import "fmt"

func main() {
		// alias string
	type NoKTP string

	var ktpNabil NoKTP = "1234567890"
	
	var contoh string = "2222222"
	var contohKTP NoKTP = NoKTP(contoh)
	
	
	fmt.Println(ktpNabil)
	fmt.Println(contohKTP)

}