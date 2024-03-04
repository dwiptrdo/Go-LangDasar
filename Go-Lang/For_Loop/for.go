package main

import "fmt"

func main() {
	// jumlah := 1

	// for jumlah <= 10 {
	// 	fmt.Println("Perulangan ke ", jumlah)
	// 	jumlah++
	// }

	for jumlah := 1; jumlah <= 10; jumlah++ {
		fmt.Println("Perulangan Ke ", jumlah)
	}
	fmt.Println("Program Selesai!")

	names := []string{"Alfredo", "Dwi", "Putra"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
