package main

import "fmt"

func main() {
	names := [...]string{"Nabil", "Argow", "Dzaky", "Rizky", "Buag", "Fora"}
	slice := names[4:6]

	fmt.Println(slice)

	slice2 := names[:3]

	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	var slice4 []string = names[:]
	fmt.Println(slice4)
	fmt.Println(append(slice4, "Nabildzr"))

	days := [...]string{
		"Senin",
		"Selasa",
		"Rabu",
		"Kamis",
		"Jumat",
		"Sabtu",
		"Minggu",
	}

	daysSlice1 := days[5:] // sabtu, minggu
	fmt.Println(daysSlice1)

	daysSlice1[0] = "Sabtu Baru"
	daysSlice1[1] = "Minggu Baru"
	fmt.Println(daysSlice1)
	fmt.Println(days)

	daysSlice2 := append(daysSlice1, "Libur Baru")
	daysSlice2[0] = "Sabtu Lama"

	fmt.Println(daysSlice1)
	fmt.Println(daysSlice2)
	fmt.Println(days)
	// array baru : "libur baru"

	var newSlice = make([]string, 2, 5)
	newSlice[0] = "Nabil"
	newSlice[1] = "Nabil"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Dzikrika")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Argow"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))	

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

}
