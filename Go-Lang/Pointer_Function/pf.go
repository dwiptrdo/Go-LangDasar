package main

import "fmt"


type Address struct {
	City, Prov, Reg string
}

func ChangeCountry(address *Address)  {
	address.Prov = "East Java"
}

func main()  {
	address := Address{}
	ChangeCountry(&address)

	fmt.Println(address)
}