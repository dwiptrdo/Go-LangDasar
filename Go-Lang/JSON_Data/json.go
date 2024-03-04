package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	FullName string `json:"Name"`
	Age int
}

func main() {
	var jsonString = `{"Name": "alfredo", "Age": 17}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user: ", data.FullName)
	fmt.Println("age: ", data.Age)
}