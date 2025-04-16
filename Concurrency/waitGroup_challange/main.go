package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMsg(s string) {
	defer wg.Done()
	msg = s
}
func printMsg() {
	fmt.Println(msg)
}

var wg sync.WaitGroup

func main() {

	//// challenge: modify this code so that the calls to updateMessage() on lines
	// 27, 30, and 33 run as goroutines, and implement wait groups so that
	// the program runs properly, and prints out three different messages.
	// Then, write a test for all three functions in this program: updateMessage(),

	msg = "hello world"

	wg.Add(1)
	go updateMsg("hello universe")
	wg.Wait()
	printMsg()

	wg.Add(1)
	go updateMsg("hello cosmos")
	wg.Wait()
	printMsg()
	
	wg.Add(1)
	go updateMsg("hello world")
	wg.Wait()
	printMsg()

}
