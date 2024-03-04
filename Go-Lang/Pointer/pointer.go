package main

import "fmt"


type Address struct {
	City, Prov, Reg string
}


func main() {
	// address1 := Address{"Blitar", "Jatim", "Indonesia"}
	// address2 := address1 // copy value
	

	// address2.Prov = "Jawa Timur"
	// fmt.Println(address1) // data tidak berubah 
	// fmt.Println(address2) // data berubah karena data di copy dari address1 
	
	// jika di pointer:
	// address1 := Address{"Blitar", "Jatim", "Indonesia"}
	// address2 := &address1 // copy value
	
	var address1 Address = Address{"Blitar", "Jatim", "Indonesia"}
	var address2 *Address = &address1 // copy value
	

	address2.Prov = "Jawa Timur"
	fmt.Println(address1) // data ikut berubah 
	fmt.Println(address2) // data berubah


}