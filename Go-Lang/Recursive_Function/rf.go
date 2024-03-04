package main

import "fmt"

func factorLoop(value int) int {
	result := 1

	for i := value; i > 0; i-- {
		result *= i
	}

	return result
}

func factorRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorRecursive(value-1)
	}

}

func main() {
	fmt.Println(factorLoop(10))
	fmt.Println(factorRecursive(10))

}
