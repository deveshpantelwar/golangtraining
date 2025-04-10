package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	//channel
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c) //go keyword for making channel
	}

	/*
		fmt.Println(<-c) // this is also a blocking call as main routine wait for a value to be returned to channel
		fmt.Println(<-c) // when more, result may be different according to the request resolving order
		fmt.Println(<-c)
		fmt.Println(<-c)
		//	fmt.Println(<-c)
		//	fmt.Println(<-c) // if waiting for more data than present than it hanags - just keep waiting to receive message
	*/

	/*
		for i := 0; i < len(links); i++ {
			fmt.Println(<-c)	// in for loop, channel is waiting to receive a messgae before moving to next iteration to receive next message
			// blocking call in for loop
		}

	*/

	//infinite loop, passing channel c instead of string as go routine

	// for {
	// 	go checkLink(<-c, c)
	// }
	//OR

	for l := range c {
		//time.Sleep(5 * time.Second)   pause main routine for 5 seconds
		//go checkLink(l,c)

		go func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}() // this whole fn is function literal -> called as soon as declared
	}

}

// fn main is main routine and fn checklink is go routine
func checkLink(link string, c chan string) {

	time.Sleep(5 * time.Second)

	//get gives 2 things response and error. we doc't care about response here
	_, err := http.Get(link) // blocking call
	if err != nil {
		fmt.Println(link, "is down!")

		//c <- "may be down" // passing msg into channel

		c <- link

		return
	}
	fmt.Println(link, "is up!")

	//c <- "yes it is up"

	c <- link
}
