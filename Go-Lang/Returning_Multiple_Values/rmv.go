package main

import "fmt"

func getFullname() (string, string) {
	return "Alfredo", "Dwi"
}

func main() {
	// firstName, midName := getFullname()
	// fmt.Println(firstName, midName)

	// menghiraukan return value dengan garis bawah
	firstName, _ := getFullname()
	fmt.Println(firstName)
}
