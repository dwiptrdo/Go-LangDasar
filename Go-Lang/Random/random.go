package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {
// 	random := rand.New(rand.NewSource(10))
// 	fmt.Println("random ke 1: ", random.Int())
// 	fmt.Println("random ke 2: ", random.Int())
// 	fmt.Println("random ke 3: ", random.Int())
// 	fmt.Println("random ke 4: ", random.Int())
// }


var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

func main() {

	fmt.Println("random string 5 karakter: ", randomString(5))


	random := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println("random ke 1: ", random.Int())
	fmt.Println("random ke 2: ", random.Float32())
	fmt.Println("random ke 3: ", random.Uint32())
	fmt.Println("random ke 4: ", random.Intn(1000))
}