package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

/*
func updateMsg(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	msg = s
	m.Unlock()
}

func main() {
	msg = "hello world"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMsg("hello universe", &mutex)
	go updateMsg("hello cosmossssssss", &mutex)
	wg.Wait()

	fmt.Println(msg)
}
*/

//for test

func updateMsg(s string) {
	defer wg.Done()

	msg = s

}

func main() {
	msg = "hello world"

	wg.Add(2)
	go updateMsg("hello universe")
	go updateMsg("hello cosmossssssss")
	wg.Wait()

	fmt.Println(msg)
}
