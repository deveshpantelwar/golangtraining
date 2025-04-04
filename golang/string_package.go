package main

import (
	"fmt"
	"strings"
)

func main() {

	data := "apple,banana,kela"
	parts := strings.Split(data, ",")
	// splits the string data on basis of seperator , here
	fmt.Println(parts)
	// op [apple banana kela]

	str := "one two three four two two five"
	count := strings.Count(str, "two")
	fmt.Println("count: ", count)

	str = "               hell!           "
	trimmed := strings.TrimSpace(str)
	fmt.Println(trimmed)

	str1 := "devesh"
	str2 := "kajla"
	//join takes 2 inputs - 1 arr of string, 2 seperator b/w them
	joined := strings.Join([]string{str1, str2}, ",")
	fmt.Println(joined)
	//devesh,kajla

	joined = strings.Join([]string{str1, "lol", str2}, " ")
	fmt.Println(joined)
	//devesh lol kajla

}
