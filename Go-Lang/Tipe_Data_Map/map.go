package main

import "fmt"


func main() {
	// var person map[string]string = map[string]string{}
	// person["nama"] = "Alfredo"
	// person["kelas"] = "XII J4"

	person := map[string]string{
		"nama":  "Alfred",
		"kelas": "J4",
	}

	fmt.Println(person["nama"])
	fmt.Println(person["kelas"])
	fmt.Println(person)

	buku := make(map[string]string)
	buku["title"] = "Buku Saya"
	buku["penulis"] = "Alfred"
	buku["pembaca"] = "Saya"

	fmt.Println(buku)

	delete(buku, "pembaca")

	fmt.Println(buku)

}
