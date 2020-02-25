package statusline

import (
	"github.com/gdamore/tcell"
	"github.com/johynpapin/nyed/ui"
)

type StatusLine struct {
	ui.Section
}

func NewStatusLine() *StatusLine {
	return &StatusLine{}
}

func (statusLine *StatusLine) Draw(screen *ui.Screen) error {
	width, height := screen.Screen.Size()
	statusLine.Width = width
	statusLine.Y = height - 1
	statusLine.Height = 1

	for x := statusLine.X; x < statusLine.X+statusLine.Width; x++ {
		for y := statusLine.Y; y < statusLine.Y+statusLine.Height; y++ {
			screen.Screen.SetContent(x, y, '-', nil, tcell.StyleDefault.Reverse(true))
		}
	}

	return nil
}
