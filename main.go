package main

import (
	"fmt"
	"math/rand"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		if b, ok := queryCache(id); ok {
			fmt.Println("From Cache:")
			fmt.Println(b)
			continue
		}
		if b, ok := queryDatabase(id); ok {
			fmt.Println(b)
			continue
		}
		fmt.Printf("Book not found with id: %v", id)
		time.Sleep(time.Second)
	}
}

func queryCache(id int) (Book, bool) {
	b, ok := cache[id]
	return b, ok
}
func queryDatabase(id int) (Book, bool) {
	time.Sleep(time.Millisecond * 100)
	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}
	return Book{}, false
}
