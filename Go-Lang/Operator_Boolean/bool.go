package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var absen = 12

	var nilaiUjian bool = nilaiAkhir > 87

	var nilaiAbsen bool = absen < 10
	var lulus bool = nilaiUjian || nilaiAbsen

	fmt.Println("Lulus = ", lulus)

	var nilaiUji = 79
	var alpha = 16

	var nilaiLulus bool = nilaiUji < 78
	var jmlAlpha bool = alpha < 10
	var lanjut bool = nilaiLulus && jmlAlpha

	fmt.Println("Lanjut Sekolah? ", lanjut)

}
