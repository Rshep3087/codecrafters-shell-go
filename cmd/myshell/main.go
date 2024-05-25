package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	repl()
}

func repl() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		input = input[:len(input)-1]
		fmt.Fprintf(os.Stdout, "%s: command not found\n", input)
	}
}
