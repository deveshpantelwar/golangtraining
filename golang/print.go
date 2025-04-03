package main

import "fmt"

func main() {

	age := 22
	name := "Devesh"
	height := 5.8234567

	fmt.Println("age:", age, "height: ", height, "name:", name)
	//age: 25height: 5.8234567name: Alice

	fmt.Println("Hello world")

	// %d integer
	// %s string
	// %c character
	// %T type of the value
	// %.3f number of float digits after dot

	fmt.Printf("age is %d\n", age)
	fmt.Printf("height is %.2f\n", height)
	fmt.Printf("Type of age is %T\n", age)
	fmt.Printf("Type of height is %T\n", height)
	fmt.Printf("name is %s\n", name)

	fmt.Printf("Age is %d\n", age)
	fmt.Printf("Name: %s, Age: %d, Height: %.2f\n", name, age, height)
}
