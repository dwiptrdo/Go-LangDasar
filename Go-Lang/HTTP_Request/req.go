package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"bytes"
)

var baseURL = "http://localhost:8080"

type student struct {
	Id    string
	Name  string
	Age int
}

// func fetchUsers() ([]student, error) {
// 	var err error
// 	var client = &http.Client{}
// 	var data []student

// 	request, err := http.NewRequest("GET", baseURL + "/users", nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response, err := client.Do(request)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer response.Body.Close()

// 	err = json.NewDecoder(response.Body).Decode(&data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return data, nil
// }

func userFetch(Id string) (student, error) {
	var err error
	var client = &http.Client{}
	var data student

	var param = url.Values{}
	param.Set("id", Id)
	var payload = bytes.NewBufferString(param.Encode())

	request, err := http.NewRequest("POST", baseURL + "/user", payload)
	if err != nil {
		return data, err
	}

	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func main()  {
	// var users, erro = fetchUsers()
	// if erro != nil {
	// 	fmt.Println("Erorr!", erro.Error())
	// 	return
	// }

	// for _, each := range users {
	// 	fmt.Printf("ID: %s\t Name: %s\t Age: %d\n", each.Id, each.Name, each.Age)
	// }

	var user1, err = userFetch("A001")
	if err != nil {
		fmt.Println("Error nih!", err.Error())
		return 
	}

	fmt.Printf("ID: %s\t Name: %s\t Age: %d\n", user1.Id, user1.Name, user1.Age)
}