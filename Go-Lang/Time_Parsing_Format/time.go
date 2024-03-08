package main

import (
	"fmt"
	"time"
)

func main()  {
	var time1 = time.Now()
	fmt.Printf("time1 %v\n", time1)

	var time2 = time.Date(2024, 3, 4, 17, 30, 0, 0, time.UTC)
	fmt.Printf("time2 %v\n", time2)

	var now = time.Now()
	fmt.Println("year:", now.Year(), "month:", now.Month(), "day:", now.Weekday().String())
}