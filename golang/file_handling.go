package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		//create a file
		file, err := os.Create("example.txt")
		if err != nil {
			fmt.Println("error while creating file :", err)
			return
		}
		defer file.Close()
		// free the resources at last

		//write contect into the file
		content := "hello world by devesh"


		_, errors := io.WriteString(file, content)
		//writestring way to write a string into a file
		// return 2 things int and error
		if errors != nil {
			fmt.Println("error while writing file :", errors)
			return
		}
		fmt.Println("successfully created file")

	*/

	/* reading from file

		file, err := os.Open("example.txt")

		if err != nil {
			fmt.Println("Error while opening file: ", err)
			defer file.Close()

		}

		defer file.Close()

		// Create a buffer to read the file content
		buffer := make([]byte, 1024)
		// Read the file content into the buffer
		for {
			n, err := file.Read(buffer)
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("error while reading file", err)
			}

			// process the read content
			fmt.Println(string(buffer[:n]))
		}

	// os.Open("example.txt") attempts to open the file named "example.txt."
	// If there's an error (e.g., the file doesn't exist), it prints an error message and exits the program.
	// defer file.Close() ensures that the file is closed when the main function exits. regardless of how it exits.
	// Buffer Creation:
	// = make([]byte, 1024) creates a byte slice (buffer) with a capacity of 1024 bytes. This buffer will be used to read chunks of the file.
	// Read File Content:
	// file.Read(buffer) reads content from the file into the buffer.
	// The loop continues until the end of the file (EOF) is reached.
	// If there's an error during the read operation (other than EOF), it prints an error message and exits the program.
	// fmt.Print(string(buffer[:n])) processes and prints the read content. In this example, it prints the content as a string.


	*/

	// Read the entire file into a byte slice
	content, err := os.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading file", err)
		return
	}
	fmt.Println(string(content))

}
