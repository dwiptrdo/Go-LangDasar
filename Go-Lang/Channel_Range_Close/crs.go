package main

import (
	"fmt"
	"runtime"
)

func send(ch chan<- string)  {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d", i)
	}
	close(ch)
}

func print(ch <- chan string)  {
	for message := range ch {
		fmt.Println(message)
	}
}

func main()  {
	runtime.GOMAXPROCS(2)

	var messages = make(chan string)
	go send(messages)
	print(messages)
}