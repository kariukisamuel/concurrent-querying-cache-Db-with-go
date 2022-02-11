package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryCache(id, m); ok {
				fmt.Println("From Cache:")
				fmt.Println(b)
			}
		}(id, wg, m)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryDatabase(id, m); ok {
				fmt.Println("From Database:")
				fmt.Println(b)
			}
		}(id, wg, m)
		time.Sleep(150 * time.Millisecond)
	}
	wg.Wait()

}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}
func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(time.Millisecond * 100)
	for _, b := range books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}
	return Book{}, false
}
