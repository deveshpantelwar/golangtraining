package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 2; i++ {
		fmt.Println("numbers are :", i)
	}

	counter := 0
	for {
		fmt.Println("infifite loop if we do not add if counter line")
		counter++
		if counter == 1 {
			break
		}
	}

	//slice
	numbers := []int{12, 45, 67}
	for index, value := range numbers {
		// range gives 2 things : index and value
		// range taking index and value from numbers
		fmt.Printf("index is %d, value is %d \n", index, value)
	}

	// range on string
	data := "DEVESH !"
	for index, char := range data {
		fmt.Printf("index is %d, character is %c\n", index, char)
	}
}
