package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func send(ch chan <- int)  {
	random := rand.New(rand.NewSource(time.Now().Unix()))

	for i := 0; true; i++ {
		ch <- i
		time.Sleep(time.Duration(random.Int()%10+1) * time.Second)
	}
}

func retreive(ch <- chan int)  {
	loop : 
	for {
		select {
		case data := <- ch:
			fmt.Print(`receive data "`, data, `"`, "\n")
		case <- time.After(time.Second * 5):
			fmt.Println("timeout. no activities under 5 second")
			break loop
		}
	}
}

func main()  {
	runtime.GOMAXPROCS(2)

	var messages = make(chan int)

	go send(messages)
	retreive(messages)
}