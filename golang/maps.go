package main

import "fmt"

func main() {

	// names <-> grades,  map is created using make method
	// map[key]value

	studentgrades := make(map[string]int)

	studentgrades["prank"] = 100
	studentgrades["khikhi"] = 10
	studentgrades["huehue"] = 40

	fmt.Println("marks of khikhi", studentgrades["khikhi"])

	studentgrades["khikhi"] = 60
	fmt.Println("new marks of khikhi", studentgrades["khikhi"])

	//deleting key value pair using key
	delete(studentgrades, "khikhi")
	fmt.Println("marks of khikhi", studentgrades["khikhi"])

	//checking if a value exists inside a map
	// when fetch data
	// 2 things return,value and bool value which tell present or no
	grades, exists := studentgrades["LolLol"]
	fmt.Println("grades of LolLol", grades)
	fmt.Println("LolLol exists :", exists)

	//for loop on map
	//printing using range bcoz range gives 2 values - index,value
	// and map also have 2 kay , value
	for index, values := range studentgrades {
		fmt.Printf("name is %s, grade is %d\n", index, values)
	}

	// unordered map so result can be in any order

	//intializing values to map while creating it
	person := map[string]int{
		"devesh": 90,
		"kajla":  80,
	}
	for index, value := range person {
		fmt.Printf("----- name is %s, marks are %d\n", index, value)
	}
}
