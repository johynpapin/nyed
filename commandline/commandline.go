package commandline

import (
	"github.com/gdamore/tcell"
	"github.com/johynpapin/nyed/screen"
	"github.com/johynpapin/nyed/ui"
)

type CommandLine struct {
	ui.Section

	CurrentCommand string
}

func NewCommandLine() *CommandLine {
	return &CommandLine{}
}

func (commandLine *CommandLine) Draw(screen *screen.Screen) error {
	if commandLine.CurrentCommand == "" {
		return nil
	}

	var x int
	for _, r := range commandLine.CurrentCommand {
		screen.Screen.SetContent(commandLine.X+x, commandLine.Y, r, nil, tcell.StyleDefault)
		x++
	}

	screen.Screen.ShowCursor(commandLine.X+x, commandLine.Y)

	return nil
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
