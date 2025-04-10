package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	//gives pointer to a file and error
	f, err := os.Open(os.Args[1])
	//Args hold the command-line arguments, starting with the program name.

	if err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
		//Exit causes the current program to exit with the given status code.
		// Conventionally, code zero indicates success, non-zero an error.
		// The program terminates immediately; deferred functions are not run.
	}

	//func io.Copy(dst io.Writer, src io.Reader) (written int64, err error)
	//Copy copies from src to dst until either EOF is reached on src or an error occurs.
	// It returns the number of bytes copied and the first error encountered while copying
	io.Copy(os.Stdout, f)
	//var os.Stdout *os.File
	//Stdin, Stdout, and Stderr are open Files pointing to the standard input, standard output, and standard error file descriptors.
}
