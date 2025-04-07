// Receiver function

package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

// create a new type of deck
//which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suits)
		}
	}

	return cards
}

//receiver is of type deck and function name is print

//any variable inside our application of type deck get access
//  to print fn

// The actual copy of the deck we're working with is available
//
//	in the function as a variable called 'd'
//
// by convention use 1-2 char of type eg d of deck
// deck as receiver fn
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// func (cards deck) print() {
// 	for i, card := range cards {
// 		fmt.Println(i, card)
// 	}
// }

// deck passed as argument in deal fn
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
	//slices from 0 to handsize and handsize to end
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	//convert deck d into string then join by seperating with ,

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	// filename of string type, data, permissoon
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename) // bs byteslice
	if err != nil {
		fmt.Println("Error :", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]

	}
}
