package main

import "fmt"

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("divisible by 0 not possible")
// 		//passing float and error value
// 		// returns a string as a value that specify error
// 	}

// 	return a / b, nil
// 	// right case, no error to pass so pass nil
// }

// passing string as error message instead Errorf
func divide(a, b float64) (float64, string) {
	if b == 0 {
		return 0, "divisible by 0 not possible"
	}

	return a / b, "nil"
}

func main() {

	//ans, err := divide(10, 2)
	// declared err so we have to use it

	// if err != nil {
	// 	fmt.Println("error handling")
	// }

	ans, _ := divide(10, 2)
	// use _ when want to ignore output
	//In the Go programming language,
	// the blank identifier, represented by an
	// underscore ( _ ), is a special identifier
	// that can be used to discard unwanted values.
	// Used when a function returns multiple values,
	// but you only need to use one or some of them.

	fmt.Println("divided value is", ans)

}
