package main

import "fmt"

func getName() (satu, dua, tiga string) {
	satu = "Alfredo"
	dua = "Dwi"
	tiga = "Putra"

	return satu, dua, tiga

}

func main() {
	a, b, c := getName()
	fmt.Println(a,b,c)

}
