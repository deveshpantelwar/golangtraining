// pointer - memory management, data sharing b/w functions
package main

import "fmt"

func modifyValueByReference(a *int) {
	*a = *a + 21
}

func main() {

	num := 2

	ptr := &num

	// var ptr *int
	// //block to which pointer points have int type value
	// ptr = &num // ptr have address of num

	fmt.Println("num :", num)
	fmt.Println("address of pointer :", ptr)
	fmt.Println("data at pointer address :", *ptr)

	var pointer *int
	fmt.Println(pointer) //op is <nil>
	// default value of pointer is nill if not assigned

	if pointer == nil {
		fmt.Println("pointer not assigned")
	}

	value := 6
	fmt.Println("value before modify :", value)
	modifyValueByReference(&value)
	fmt.Println("value after modify :", value)

}
