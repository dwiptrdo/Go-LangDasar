package main

import (
	"go-lang/helper"
	"go-lang/database"
	_"go-lang/internal"
	"fmt"
)

func main() {
	result := helper.SayHalo("Alfredo")
	fmt.Println(result)


	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak bisa diakses hanya bisa digunakan di package yang asli
	// fmt.Println(helper.sayHaii) // tidak bisa diakses


	// database
	fmt.Println(database.GetDatabase())
}