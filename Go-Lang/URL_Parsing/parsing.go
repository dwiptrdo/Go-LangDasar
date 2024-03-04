package main

import (
	"fmt"
	"net/url"
)

func main() {
	var urlString = "https://et.com:90/haii?name=alfredo&age=17"
	var u, e = url.Parse(urlString)
	if e != nil {
		fmt.Println(e.Error())
		return
	}

	fmt.Printf("url: %s\n", urlString)

	fmt.Printf("protocol: %s\n", u.Scheme) // https
	fmt.Printf("host: %s\n", u.Host) // et.com:90
	fmt.Printf("path: %s\n", u.Path) // /haii

	var name = u.Query()["name"][0] // alfredo
	var age = u.Query()["age"][0] // 17
	fmt.Printf("name: %s, age: %s", name, age)

}