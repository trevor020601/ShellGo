package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	cmdReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")

		// Reads from the keyboard
		cmdInput, e := cmdReader.ReadString('\n')
		if e != nil {
			fmt.Fprintln(os.Stderr, e)
		}

		if e = execInput(cmdInput); e != nil {
			fmt.Fprintln(os.Stderr, e)
		}
	}
}

func execInput(input string) error {
	// Removes newline
	input = strings.TrimSuffix(input, "\n")

	args := strings.Split(input, " ")

	switch args[0] {
	case "exit":
		os.Exit(0)
	case "cd":
		if len(args) < 2 {
			return errors.New("path required")
		}
		return os.Chdir(args[1])
	}

	// Fix for Windows?
	// need to add /c somewhere I think
	cmd := exec.Command("cmd.exe", args...)

	//cmd := exec.Command(args[0], args[1:]...)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
