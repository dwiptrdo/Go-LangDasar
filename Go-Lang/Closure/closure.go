package main

import "fmt"

func main() {
	counter := 0

	imcrement := func() {
		fmt.Println("Increment")
		counter++
	}

	imcrement()
	imcrement()
	imcrement()

	fmt.Println(counter)
}
