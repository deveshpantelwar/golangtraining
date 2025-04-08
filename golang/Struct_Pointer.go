package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jim.updateName("jimmy")
	jim.print()

	jimPointer := &jim
	jimPointer.updateNamePointer("jimmy")

	//OR jim.updateNamePointer("jimmy")
	//go shortcut for pointer , no need to point to address
	//can directly use pointer operator in func
	//func (pointerToPerson *person) updateNamePointer(
	jim.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

//TURN ADDRESS INTO VALUE WITH *ADDRESS
//TURN VALUE INTO ADDRESS WITH &VALUE

func (pointerToPerson *person) updateNamePointer(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

/*This is a type description -it means we're working with a pointer to a person

func (pointer ToPerson *person) updateName() {
	*pointer To Person
}
This is an operator - it means we want to manipulate the value the pointer is referencing

*/
