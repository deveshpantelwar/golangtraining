package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {

	t := triangle{base: 10, height: 10}
	s := square{sideLength: 10}

	printArea(t)
	printArea(s)

}

func printArea(s shape) {
	fmt.Println("area is :", s.getArea())

}

// languagebot didn't have any argument in struct so only wrote receiver fn name eg triangle not t triangle
// to use fields of triangle use t.base
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
