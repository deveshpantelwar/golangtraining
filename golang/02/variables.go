package main

//package main - main is used to build executable files

import (
	"fmt"
)

func main() {

	fmt.Println("let's learn go ")
	var name string = "devesh"

	surname := "kajla"
	// no need to mention var and datatype.
	// := works only inside a function , not in package area.

	fmt.Println(name)
	fmt.Println(surname)

	var version = 76

	fmt.Println(version)
	var money int = 67000

	var currency = 3489

	fmt.Println(money)

	fmt.Println("currency: ", currency)

	var dimension float64 = 87.12
	fmt.Println(dimension)

	var decided bool = false
	decided = true
	// can change value of variables declared as var
	fmt.Println(decided)

	const pi = 3.14
	//pi = 4.14
	// cannot cahnge value as declared as constant variable

	//Println is function is written in some fmt package which we are using in main
	var Public = "data is important"
	var private = "data is private"
	fmt.Println(Public)
	fmt.Println(private)

	//In this example, PublicVariable is visible and can be accessed from other packages,
	//while privateVariable is only visible within the same package.

}
