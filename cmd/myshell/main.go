package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	repl()
}

// repl stands for Read-Eval-Print-Loop
// It reads input from the user, evaluates it, prints the result, and loops back to the start
func repl() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd := parse(input)
		execute(cmd)
	}
}

// command represents a command that the user wants to execute
type command struct {
	name string
	args []string
}

// parse takes a string and returns a command
func parse(input string) command {
	cmds := strings.Fields(input)
	return command{cmds[0], cmds[1:]}
}

// execute takes a command and executes it
func execute(cmd command) {
	switch cmd.name {
	case "exit":
		// get the exit code from the first argument
		code, _ := strconv.Atoi(cmd.args[0])
		os.Exit(code)
	default:
		fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd.name)
	}
}
