package main

import (
	"bufio"
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

	cmd := exec.Command(input)

	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd.Run()
}
