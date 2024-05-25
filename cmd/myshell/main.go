package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// strip the return character and newline
	input = input[:len(input)-1]

	// Print the input
	fmt.Fprintf(os.Stdout, "%s: command not found\n", input)
}
