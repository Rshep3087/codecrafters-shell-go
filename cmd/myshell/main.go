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

func repl() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		cmd := parse(input)
		execute(cmd)
	}
}

type command struct {
	name string
	args []string
}

func parse(input string) command {
	cmds := strings.Fields(input)
	return command{cmds[0], cmds[1:]}
}

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
