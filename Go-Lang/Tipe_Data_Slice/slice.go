package main

import "fmt"

func main() {
	names := [...]string{"Alfredo", "Dwi", "Putra", "Aziz", "Munib", "Ali", "Fajar"} // jika dijabarkan maka 'var values = [...]string{..., ..., ...}'
	slice1 := names[3:6]                                                             // jika dijabarkan maka 'var slice []string = names[3:6]'

	fmt.Println(slice1)

	slice2 := names[:3]

	fmt.Println(slice2)

	slice3 := names[3:]

	fmt.Println(slice3)

	slice4 := names[:]

	fmt.Println(slice4)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}
	daySlice1 := days[5:] // Sabtu Minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Libur"
	daySlice1[1] = "Minggu Libur"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Hari Libur")
	daySlice2[0] = "Sabtu Lama"
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)

	var newSlice1 []string = make([]string, 2, 5)
	newSlice1[0] = "Alfredo"
	newSlice1[1] = "Dwi"
	// newSlice[2] = "Putra" // error harusnya menggunakan append

	fmt.Println(newSlice1)
	fmt.Println(len(newSlice1))
	fmt.Println(cap(newSlice1))

	newSlice2 := append(newSlice1, "Dodo")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Putra"
	fmt.Println(newSlice2)
	fmt.Println(newSlice1)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)

}
