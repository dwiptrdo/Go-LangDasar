package main

import "fmt"

func main()  {
	var (
		l = 17
		m = 15
		k = l + m
	)

	fmt.Println(k)

	var (
		q = 17
		w = 15
		e = q - w
	)

	fmt.Println(e)

	var (
		r = 17
		t = 15
		y = r * t
	)

	fmt.Println(y)

	var (
		u = 17
		i = 15
		o = u / i
	)

	fmt.Println(o)

	var (
		p = 17
		a = 15
		b = p % a
	)

	fmt.Println(b)

	// Augmented 
	var z = 12
	z += 3

	fmt.Println(z)

	// Unary Operator
	var x = 2
	x++ 
	x++

	fmt.Println(x)

	
}