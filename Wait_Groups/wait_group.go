package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{"alpha", "beta", "gamma", "delta", "theta"}
	//	go printSomething("this is 1st string")

	/*
		for i, x := range words {
			go printSomething(fmt.Sprintf("%d : %s", i, x))
			// go printSomething(fmt.Printf("%v : ",i))
			//replace with 	go printSomething(fmt.Sprintf("%v : ", i))

			// Problem Breakdown:
			// fmt.Printf(...) directly terminal pe print karta hai aur (int, error) return karta hai, na ki string.
			// Tum us return value ko printSomething(...) ko de rahe ho, jo ek string expect karta hai, isiliye compiler bol raha hai:

			//📌 Summary:

			// Function	Print	 karta hai?		String return karta hai?			Use-case
			// fmt.Printf		  ✅ Yes			❌ No						Direct terminal output
			// fmt.Sprintf	    	❌ No		✅ Yes					String ko store/process karna

		}
		time.Sleep(1 * time.Second)
		printSomething("this is 2nd string ")
	*/

	wg.Add(5) // 5 length of slice words

	for i, x := range words {
		go printSomething(fmt.Sprintf("%v : %v", i, x), &wg)
	}
	wg.Wait()
	// wg.Add(1)
	//printSomething("this is the 2nd thing to print!", &wg)
}
