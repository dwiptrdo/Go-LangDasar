package main

import "fmt"

func logging() {
	fmt.Println("Fungsi Selesai!")
}

func runApp() {
	defer logging()
	fmt.Println("Run Aplication")
}

func main() {
	runApp()
}
