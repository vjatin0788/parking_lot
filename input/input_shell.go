package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Shell struct {
	Cmd string
}

//Create the interactive shell to execute the commands.
func CreateIntractiveShell() *Shell {
	return &Shell{
		Cmd: "$ ",
	}
}

func (s *Shell) RunShell() {
	scanner := bufio.NewScanner(os.Stdin)
	s.newLine()
	for scanner.Scan() {
		command := strings.ToLower(scanner.Text())
		err := processCommands(strings.Fields(command))
		if err != nil {
			fmt.Println(err)
		}
		s.newLine()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func (s *Shell) newLine() {
	fmt.Print(s.Cmd)
}
