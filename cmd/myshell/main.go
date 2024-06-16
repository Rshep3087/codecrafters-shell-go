package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"slices"
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
		if len(cmd.args) == 0 {
			os.Exit(0)
		}
		code, _ := strconv.Atoi(cmd.args[0])
		os.Exit(code)
	case "echo":
		args := strings.Join(cmd.args, " ")
		fmt.Println(args)
		return
	case "type":
		handleType(cmd)
		return
	case "pwd":
		dir, _ := os.Getwd()
		fmt.Println(dir)
		return
	case "cd":
		handleCD(cmd)
		return
	}

	excmd := exec.Command(cmd.name, cmd.args...)
	excmd.Stdout = os.Stdout
	excmd.Stderr = os.Stderr
	err := excmd.Run()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s: command not found\n", cmd.name)
		return
	}
}

func handleCD(cmd command) {
	dir := cmd.args[0]

	err := os.Chdir(dir)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Printf("%s: No such file or directory\n", dir)
			return
		}
	}
}

func handleType(cmd command) {
	if slices.Contains([]string{"exit", "echo", "type"}, cmd.args[0]) {
		fmt.Fprintf(os.Stdout, "%s is a shell builtin\n", cmd.args[0])
		return
	}

	if fp, ok := inPath(cmd.args[0]); ok {
		fmt.Fprintf(os.Stdout, "%s\n", fp)
		return
	}

	fmt.Fprintf(os.Stdout, "%s: not found\n", cmd.args[0])
}

func inPath(t string) (string, bool) {
	path := os.Getenv("PATH")

	dirs := strings.Split(path, ":")
	for _, dir := range dirs {
		files, _ := os.ReadDir(dir)
		for _, file := range files {
			if file.Name() == t {
				return fmt.Sprintf("%s/%s", dir, file.Name()), true
			}
		}
	}

	return "", false
}
