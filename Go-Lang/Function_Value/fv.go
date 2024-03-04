package main

import "fmt"

func bye(name string) string {
	return "Bye " + name
}

func main()  {
	dahh := bye
	fmt.Println(dahh("Dodo"))
}
