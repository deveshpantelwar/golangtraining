package main

import (
	"math/rand"
	"time"
)


const NumberOfPizza = 10

var pizzasMade, pizzaFailed, total int

type Producer struct{
	data chan PizzaOrder
	quit chan chan error
}

type PizzaOrder struct{
	pizzaNumber int
	message string
	success bool
}


func made(){

	rand.Seed(time.Now().UnixNano())

	

}