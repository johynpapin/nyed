package statusline

import (
	"github.com/gdamore/tcell"
	"github.com/johynpapin/nyed/screen"
	"github.com/johynpapin/nyed/ui"
	"github.com/johynpapin/nyed/utils"
)

type StatusLine struct {
	ui.Section

	CurrentMode utils.Mode
}

func NewStatusLine() *StatusLine {
	return &StatusLine{
		CurrentMode: utils.MODE_NORMAL,
	}
}

func (statusLine *StatusLine) Draw(screen *screen.Screen) error {
	modeStyle := utils.GetModeStyle(statusLine.CurrentMode)
	modeText := " " + utils.GetModeText(statusLine.CurrentMode) + " "

	var x int
	for _, r := range modeText {
		screen.Screen.SetContent(statusLine.X+x, statusLine.Y, r, nil, modeStyle)
		x++
	}

	for ; x < statusLine.Width; x++ {
		screen.Screen.SetContent(statusLine.X+x, statusLine.Y, ' ', nil, tcell.StyleDefault.Background(tcell.ColorLightSlateGrey))
	}

	return nil
}
