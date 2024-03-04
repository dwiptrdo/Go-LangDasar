package main

import "fmt"

type Customer struct {
	name, country string
	age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Halo", name, "nama saya ", customer.name)
}

func main() {
	var Dodo Customer

	Dodo.name = "Alfredo Dwi Putra"
	Dodo.country = "Indonesia"
	Dodo.age = 17

	fmt.Println(Dodo)
	fmt.Println(Dodo.name)
	fmt.Println(Dodo.country)
	fmt.Println(Dodo.age)

	Muner := Customer{
		name:    "Munib F.A",
		country: "Indonesia",
		age:     18,
	}

	fmt.Println(Muner)

	Ali := Customer{"Fajar A.R", "Indonesia", 18}
	fmt.Println(Ali)

	Ali.sayHello("Rohman")

	Dodo.sayHello("Rohman")

}
