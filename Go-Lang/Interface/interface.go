package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(value HasName) {
	fmt.Println("Haloo", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Nama string
}

func (animal Animal) GetName() string {
	return animal.Nama

}

func main() {
	person := Person{Name: "Putra"}
	sayHello(person)

	animal := Animal{Nama: "Burung"}
	sayHello(animal)
}
