package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num int = 42
	fmt.Println("Number is", num)
	fmt.Printf("Type of num is %T\n", num)

	var data float64 = float64(num)
	fmt.Println("data is :", data)
	fmt.Printf("Type of data is %T\n", data)

	// int to string
	num = 123
	str := strconv.Itoa(num)
	fmt.Println("str is ", str)
	fmt.Printf("Type of str is %T\n", str)

	// string to int
	number_string := "1234"
	number_int, _ := strconv.Atoi(number_string)
	//atoi returns 2 things, 1 is integer, 2 is error handled by _
	fmt.Println("number_int is", number_int)
	fmt.Printf("Type of number_int is %T\n", number_int)

	// string to 64 bit float i.e. float64
	num_string := "3.14"

	number_float, _ := strconv.ParseFloat(num_string, 64)
	fmt.Println("number_float is", number_float)
	fmt.Printf("Type of number Loat is %T\n", number_float)


	//float to string
	s := fmt.Sprintf("%f", 123.456) // s == "123.456000"
	fmt.Println("string of float is :", s)
	fmt.Printf("type is %T\n", s)

}
