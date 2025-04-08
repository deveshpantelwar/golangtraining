package main

import "fmt"

func main() {
	day := 2
	switch day {
	case 1, 30, 22: // multiple value check
		fmt.Println("good day")
	case 2: // colon important
		fmt.Println("bad day")
	default:
		fmt.Println("lucky day")
	}

	month := " december"
	switch month {
	case "January", "February", "March":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Spring")
	default:
		fmt.Println("Other season")
	}

	temperature := -10
	switch  {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature >= 0 && temperature < 10:
		fmt.Println("Cold")
	case temperature >= 10 && temperature < 20:
		fmt.Println("Cool")
	case temperature >= 20 && temperature < 30:
		fmt.Println("Warm")
	default:
		fmt.Println("hot")
	}
}
