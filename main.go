package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	terminal := tview.NewApplication()
	//text := tview.NewInputField().SetLabel("> ").SetLabelColor(tcell.ColorDarkGreen)
	background := tview.NewBox().SetBackgroundColor(tcell.ColorBlack).SetBorder(true).SetTitle("ShellGo").SetTitleColor(tcell.ColorDarkGreen)

	if err := terminal.SetRoot(background, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}

	terminal.SetInputCapture(func(input *tcell.EventKey) *tcell.EventKey {
		if input.Rune() == 113 {
			terminal.Stop()
		}
		return input
	})

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
