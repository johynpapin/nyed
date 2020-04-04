package state

import (
	"github.com/gdamore/tcell"
)

type CommandLine struct {
	Width, Y int

	CurrentCommand string
}

func NewCommandLine() *CommandLine {
	return &CommandLine{}
}

func (commandLine *CommandLine) HandleEventKey(eventKey *tcell.EventKey) (string, error) {
	if eventKey.Key() == tcell.KeyRune {
		commandLine.CurrentCommand += string(eventKey.Rune())
		return "", nil
	}

	if eventKey.Key() == tcell.KeyEscape {
		commandLine.CurrentCommand = ""
		return "", nil
	}

	if eventKey.Key() == tcell.KeyEnter {
		command := commandLine.CurrentCommand
		commandLine.CurrentCommand = ""
		return command[1:], nil
	}

	if eventKey.Key() == tcell.KeyBackspace2 {
		commandLine.CurrentCommand = commandLine.CurrentCommand[:len(commandLine.CurrentCommand)-1]
		return "", nil
	}

	return "", nil
}
