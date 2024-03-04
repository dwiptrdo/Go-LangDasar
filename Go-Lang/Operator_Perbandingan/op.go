package main

import "fmt"

func main() {
	var name1 = "Alfredo"
	var name2 = "Alfredo"

	var result bool = name1 == name2

	fmt.Println(result)

	var name3 = "Alfredo"
	var name4 = "Dwi"

	var results bool = name4 != name3

	fmt.Println(results)

}