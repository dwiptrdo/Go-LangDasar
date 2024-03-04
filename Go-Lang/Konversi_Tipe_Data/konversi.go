package main

import "fmt"

func main() {
	// Konversi Tipe Data INT
	var nilai32 int32 = 32853
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32) // hasilnya tidak diterima karena melebihi kapasitas, maka hasilnya pasti akan minus(-)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	// Konversi Tipe Data String
	var nama = "Alfredo Dwi Putra"
	var a = nama[0]
	var aNama = string(a)

	fmt.Println(nama)
	fmt.Println(a)
	fmt.Println(aNama)
}
