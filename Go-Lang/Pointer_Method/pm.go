package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main()  {
	dodo := Man{"Dodo"}
	dodo.Married()

	fmt.Println(dodo.Name)
}