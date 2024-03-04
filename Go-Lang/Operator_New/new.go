package main

import "fmt"


type Address struct {
	City, Prov, Reg string
}

func main() {
	var alamat1 *Address = &Address{}
	var alamat2 *Address = alamat1

	alamat2.City = "Tangerang"

	fmt.Println(alamat1)
	fmt.Println(alamat2)

}