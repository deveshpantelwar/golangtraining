package main

func main() {

	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	//calling print fn from deck instead of writing it

	// 	cards := newDeck()

	// 	//deal return two values of type deck and input deck, int
	// 	hand, remainingCards := deal(cards, 5)
	// 	hand.print()
	// 	remainingCards.print()
	//

	// hello := "hi there"
	// fmt.Println([]byte(hello)) // string into bytle slices

	// cards := newDeck()
	// fmt.Println(cards.toString())

	// cards := newDeck()
	// cards.saveToFile("my_cards")
	//created a new file my_cards

	/*	cards := newDeckFromFile("my_cards")
		cards.print()
	*/
	// if acces a file that doesn't exist
	//Error : open my_card: no such file or directory
	//exit status 1

	cards := newDeck()
	cards.shuffle()
	cards.print()

}
