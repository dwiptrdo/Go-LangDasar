package main

import "fmt"

func endApp() {
	fmt.Println("End Aplication")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Yah ERROR")
	}
}

func main() {
	runApp(true)
	fmt.Println("Alfredo dwi Putra")
}
