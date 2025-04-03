package main

import "fmt"

func main() {
	x := 6
	if x > 5 {
		fmt.Println(x)
	}

	z := 10
	if z > 10 {
		fmt.Println("z is greater than 10")
	} else if z > 5 {
		fmt.Println("z is greater than 5 but smaller than 10")
	} else {

		fmt.Println("z is smaller than 5")
	}
	y := 10
	if x < 5 && (y > 5 || z < 43) {
		//go doesn't remove paranthesis to prevent logic here
		fmt.Println("hey how are you?")
	} else {
		fmt.Println("Learn programming with Hello World")
	}
}
