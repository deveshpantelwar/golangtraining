package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Println("Learning web requests")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting GET response ", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Type of response %T \n", res)

	//read the response
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("error reading response", err)
		return
	}

	fmt.Println("response ", string(data))
	//This line converts the byte slice body to a string and
	// prints it to the console. The response body typically
	// contains the content (data) returned by the server in
	// response to the request.

}
