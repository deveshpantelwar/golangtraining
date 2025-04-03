//In Go, a slice is a flexible and dynamic data structure that provides a more powerful alternative to arrays.
//Unlike arrays, slices have a dynamic size, and their length can be changed during the program's execution.

package main

import "fmt"

func main() {
	//initialization
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6, 7, 8, 99)
	fmt.Println("Numbers :", numbers)
	fmt.Println("lenght :", len(numbers))
	fmt.Printf("data type : %T\n", numbers)
	//slice is based on array so data type is array of int

	name := []string{} // empty slice
	fmt.Println("name :", name)

	values := make([]int, 0, 5)
	// slice of int array with length 3, capacity 5
	values = append(values, 4)
	values = append(values, 4)
	// now slice is 0 0 0 4 4, length 5, capacity 5

	values = append(values, 5)
	//now slice 0 0 0 4 4 5, length 6, capacity 10
	// capacity doubles when filled like vector

	fmt.Println("Values :", values)
	fmt.Println("Length :", len(values))
	fmt.Println("capacity :", cap(values))
}
