package main

import "fmt"

type bot interface {
	getGreeting() string
	//if you are a type in this program with a function called
	//getGreeting() and you return a string then you are now
	//an honorary member of typr 'bot'

	//or any type in program with a getGreeting() and return string
	// are included as type bot
}

type englishBot struct{}//becomes also of type bot
type spanishBot struct{}
//now that you are also an member of type 'bot', you can now
//call this function called 'printGreeting()

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

}

// func printGreeting(eb englishBot){
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot){
// 	fmt.Println(sb.getGreeting())
// }

func printGreeting(b bot) {
	//it only accept value of type bot but eb and sb are
	//  pretend bot now so acceptable
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//getGreeting expect to have englishBot as receiver
	return "hello!"
}

func (spanishBot) getGreeting() string {
	return "hola!"
}
