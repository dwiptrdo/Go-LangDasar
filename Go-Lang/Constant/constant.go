package main

import "fmt"

func main()  {
	const firstName string = "Alfredo"
	const midName = "Dwi"
	const lastName = "Putra"

	const (
		nama = "Alfredo Dwi Putra"
		panggilan = "dodo"

	)

	fmt.Println(firstName, midName, lastName)
	fmt.Println(nama, panggilan)

	// Constan tidak bisa diubah, jika memaksa diubah maka akan error
	// firstName = "Dwi"
	// midName = "Putra"

	
}