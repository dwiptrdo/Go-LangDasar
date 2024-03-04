package main

import "fmt"

func main() {
	nama := "Alfredo"

	if nama == "Dodo" {
		fmt.Println("Haloo Dodo")
	} else if nama == "Dwi" {
		fmt.Println("Halo Diriku!")
	} else {
		fmt.Println("Kamu siapa?")
	}

	
	if length := len(nama) ; length > 5 {
		fmt.Println("Nama Terlalu panjang")
	} else {
		fmt.Println("Nama benar")
	}
}
