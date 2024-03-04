package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHello = func (who string)  {
		var data = fmt.Sprintf("haii %s", who)
		messages <- data
	}

	go sayHello("Alfredo")
	go sayHello("Dwi")
	go sayHello("Putra")

	var message1 = <-messages
	fmt.Println(message1)
	
	var message2 = <-messages
	fmt.Println(message2)

	var message3 = <-messages
	fmt.Println(message3)
}