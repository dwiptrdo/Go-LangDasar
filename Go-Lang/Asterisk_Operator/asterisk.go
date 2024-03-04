package main

import "fmt"

type Address struct {
	City, Prov, Reg string
}

func main() {
	var address1 Address = Address{"Blitar", "Jatim", "Indonesia"}
	var address2 *Address = &address1 // pointer
	

	address2.Prov = "Jawa Timur"
	fmt.Println(address1) // data ikut berubah 
	fmt.Println(address2) // data berubah menjadi Jawa Timur

	// address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
}