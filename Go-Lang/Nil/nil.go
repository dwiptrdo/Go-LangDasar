package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"nane": name,
		}
	}
}

func main() {
	data := NewMap("P")

	if data == nil {
		fmt.Println("data map masih kosong")
	} else {
		fmt.Println(("datanya ada"))
	}

	
}