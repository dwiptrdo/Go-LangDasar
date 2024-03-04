package main

import "fmt"

func main() {
	type noHp string

	var Hp noHp = "08129999999"

	fmt.Println(Hp)
	fmt.Println(noHp("00891222222"))

	var Hp2 string = "91110000010"
	var contohHp2 noHp = noHp(Hp2)

	fmt.Println(contohHp2)

}
