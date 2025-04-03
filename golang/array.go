package main

import "fmt"

func main() {
	var name [5]string
	//giving size is important
	name[0] = "devesh"
	name[1] = "aman"

	fmt.Println("names without quote", name)
	// names without quote [devesh aman   ]

	fmt.Printf("names are %q \n", name)
	//quotted string, nil values of name value will be shown in quotes ""
	// names are ["devesh" "aman" "" "" ""]

	//var numbers = [5]int{1, 2, 3, 4, 5}
	numbers := [8]int{1, 2, 3, 4}
	//initialising elements while declaring arrayl

	fmt.Println("numbers are", numbers)
	fmt.Println("length of numbers is", len(numbers))
	fmt.Println("value at 2nd index", name[1])
}
