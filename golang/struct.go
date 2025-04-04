package main

import "fmt"

// data types int, string, person, contact, address
// var name string
type person struct {
	Firstname string
	Lastname  string
	Age       int
}

type contact struct {
	Email string
	Phone string
	fax   string
}

type address struct {
	House int
	Area  string
	State string
}

type employee struct {
	//variable_name , data type like always
	person_details person
	person_contact contact
	person_address address
}

func main() {

	//var name string

	var devesh person
	fmt.Println("Data of devesh :", devesh)
	// op """"0 // empty string nil , nill, 0// __0

	// declaring variable and filling values
	devesh.Firstname = "devesh"
	devesh.Lastname = "kajla"
	devesh.Age = 22
	//fmt.Println("Data of devesh :", devesh)

	//2nd method initialize while declaring
	person1 := person{
		Firstname: "hello",
		Lastname:  "world",
		Age:       9999999,
	}
	fmt.Println("person 1 :", person1)
	//person 1 : {hello world 9999999}

	//3rd method using new keyword
	var person2 = new(person)
	person2.Firstname = "dom"
	person2.Lastname = "pandit"
	person2.Age = 59

	//fmt.Println("person 2 :", person2)
	// person 2 : &{dom pandit 59}
	// new keyword makes person2 a pointer so &

	var employee1 employee
	employee1.person_details = person{
		Firstname: "mukes",
		Lastname:  "ambani",
		Age:       55,
	}

	employee1.person_contact.Email = "hi@gmail.com"
	employee1.person_contact.Phone = "9876543"

	employee1.person_address = address{
		House: 504,
		Area:  "Rohini",
		State: "Delhi",
	}

	fmt.Println(employee1)
	fmt.Println(employee1.person_address)
	fmt.Println(employee1.person_address.State)

	employee1.person_contact.fax = "fax1234"
	fmt.Println("updated fax number later ", employee1.person_contact)

}
