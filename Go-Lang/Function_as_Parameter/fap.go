package main

import "fmt"

type Filter func(string) string

// func withFilter(nama string, filter func(string) string) {
// 	filterName := filter(nama)
// 	fmt.Println("Halo " + filterName)
// }

func withFilter(nama string, filter Filter) {
	filterName := filter(nama)
	fmt.Println("Halo " + filterName)
}

func spamFilter(nama string) string {
	if nama == "Anjing" {
		return "SENSOR!!"
	} else {
		return nama
	}

}

func main() {
	withFilter("Alfredo", spamFilter)

	filter := spamFilter
	withFilter("Anjing", filter)

}
