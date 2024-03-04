package main

import "fmt"

func sayHelloTo(firstName string, midName string) {
	fmt.Println("Hello", firstName, midName)
}


func main()  {
	sayHelloTo("Alfredo", "Dwi")
	sayHelloTo("Dwi", "Putra")
}