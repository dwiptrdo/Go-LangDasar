package main

import "fmt"

func main()  {
	var saya = "Alfredo Dwi Putra"
	fmt.Println(saya)

	var nama string

	nama = "Alfredo"
	fmt.Println(nama)

	nama = "Alfredo Dwi"
	fmt.Println(nama)

	// bisa menggunakan cara cepat untuk membuat variable, tapi hanya untuk deklarasi pertama
	// nama := "Putra"
	// fmt.Println(nama)

	var (
		firstName = "Alfredo"
		midName = "Dwi"
		lastName = "Putra"
	)

	fmt.Println(firstName)
	fmt.Println(midName)
	fmt.Println(lastName)
	
}