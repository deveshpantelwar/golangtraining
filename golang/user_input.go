package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hey , what is yor name")
	// var name string
	// fmt.Scan(&name)
	// fmt.Println("so your name is", name)

	// Scan breaks the input after it encounter space

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("your name is", name)

	// bufio. NewReader (os. Stdin) creates a new buffered reader that reads from the standard input (os. Stdin).

	// reader.ReadString('\n') reads a line from the input until it encounters a newline character ('\n').
	// This allows the program to read the entire line of input, including spaces.

	// The input is stored in the name variable, and any potential errors are handled.

	// A buffered reader is a type of reader that reads data from an underlying source,
	// such as a file or standard input (keyboard), and stores that data in a buffer. The buffer is a temporary storage area in memory.
	// Buffered readers are commonly used to improve the efficiency of input operations.
}
