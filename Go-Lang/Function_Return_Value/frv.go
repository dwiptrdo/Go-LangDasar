package main

import "fmt"

func getHello(name string) string  {
	hello := "Halo " + name
	return hello
}


func main()  {
	result := getHello("Alfredo")
	fmt.Println(result)

	fmt.Println(getHello("Dodo"))
}