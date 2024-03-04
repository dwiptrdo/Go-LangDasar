package main

import "fmt"

func main() {
	name := "Dodo"

	switch name {
	case "Dodo":
		fmt.Println("Haii!")

	case "Alfredo":
		fmt.Println("Haloo!")

	default:
		fmt.Println("Kamu Siapa?")
	}

	switch length := len(name); length < 5 {
	case true:
		fmt.Println("Sudah Benar")
	case false:
		fmt.Println("Nama masih salah")
	}


	name = "Alfredo"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Pendeknya")
	case length > 5:
		fmt.Println("Panjangnya")
	default:
		fmt.Println("Nama sudah benar")
	}

}
