package helper

import "fmt"

var version = "1.0.1"
var Application = "Go-Lang"

func sayHai(name string) string {
	return "Haii " + name
}

func SayHalo(name string) string {
	return "Halo " + name
}

func Contoh()  {
	sayHai("All")
	fmt.Println(version)
}