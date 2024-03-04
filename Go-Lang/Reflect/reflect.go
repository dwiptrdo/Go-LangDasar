package main

import (
	"fmt"
	"reflect"
)

// Mencari Tipe Data & Value Menggunakan Reflect

func main() {
	var angka = "haii"
	var reflectValue = reflect.ValueOf(angka)

	fmt.Println("tipe Variabel", reflectValue.Type())
	fmt.Println("nilai Variabel", reflectValue.Interface()) // Jika nilai hanya diperlukan untuk ditampilkan ke output, bisa menggunakan Interface

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variabel", reflectValue.Int())
	} else {
		fmt.Println("Salah tipe data!")
	}
}
