package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Alfredo"
	name[1] = "Dwi"
	name[2] = "Putra"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [...]int{ // gunakan titik sebanyak 3 kali jika jumlahnya tidak mau ditentukan di awal
		90,
		80,
		85,
	}

	fmt.Println(values)
}
