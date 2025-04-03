package main

import "fmt"

func simplefn() {
	fmt.Println("simple function")
}

//func add(a, b int) (int) {
//return a+b}
//func add(a int, b int) (int)
//func add(a, c, d, b int) (int)

//a and b are input parameters, (int) is datatype of return
//if only 1 return type , no need for paranthesis ()

func add(a, b int) (result int) {
	// return variable result of int type
	result = a + b
	return
	//or return a+b will also return sum in result

}

func multiply(a, b int) (result int) {
	result = a * b
	return
}

func main() {

	fmt.Println("functions in golang")
	simplefn()
	ans := add(2, 7)
	fmt.Println("addition is", ans)

	multi := multiply(3, 4)
	fmt.Println("multiplication is ", multi)
}
