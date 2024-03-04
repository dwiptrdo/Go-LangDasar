package main

import "fmt"

func getSum(numbers ...int) int {
	jumlah := 0

	for _, number := range numbers {
		jumlah += number
	}

	return jumlah
}

func main() {
	fmt.Println(getSum(10,1,19))
	fmt.Println(getSum(10,1,19,10,7,3,17))

	numbers := []int {10,10,10,10}
	fmt.Println(getSum(numbers...))

}
