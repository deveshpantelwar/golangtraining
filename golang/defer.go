package main

import "fmt"

func add(a, b int) int {
	return a + b
}
func main() {

	fmt.Println("Starting of the program")

	data := add(5, 6)
	defer fmt.Println("Data is ", data)
	defer fmt.Println("Middle of the program")
	fmt.Println("End of the program")

	// defer makes tha statement execute before ending of main
	//execute multiple defer in LIFO order

	//starting of the program
	//end of the program
	// Middle of the program
	// Data is:", 11

}
