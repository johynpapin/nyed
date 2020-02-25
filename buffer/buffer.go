package buffer

import (
	"github.com/gdamore/tcell"
	"github.com/johynpapin/nyed/ui"
	"unicode/utf8"
)

type Buffer struct {
	ui.Section

	Lines            []string
	currentLine      int
	CursorX, CursorY int
}

func NewBuffer() *Buffer {
	return &Buffer{
		Lines: []string{""},
	}
}

func (buffer *Buffer) Draw(screen *ui.Screen) error {
	width, height := screen.Screen.Size()
	buffer.Width = width
	buffer.Height = height - 1

	xLimit := buffer.X + buffer.Width
	yLimit := buffer.Y + buffer.Height

	if len(buffer.Lines) < yLimit {
		yLimit = len(buffer.Lines)
	}

	for y := buffer.Y; y < yLimit; y++ {
		x := 0
		for _, r := range buffer.Lines[y] {
			if x >= xLimit {
				break
			}

			screen.Screen.SetContent(x, y, r, nil, tcell.StyleDefault)
			x++
		}

		for ; x < xLimit; x++ {
			screen.Screen.SetContent(x, y, ' ', nil, tcell.StyleDefault)
		}
	}

	screen.Screen.ShowCursor(buffer.CursorX, buffer.CursorY)

	return nil
}

func (buffer *Buffer) HandleEventKey(eventKey *tcell.EventKey) error {
	if eventKey.Key() == tcell.KeyRune {
		buffer.Lines[buffer.currentLine] += string(eventKey.Rune())
		buffer.CursorX++
	}

	if eventKey.Key() == tcell.KeyEnter {
		buffer.currentLine++
		if buffer.currentLine == len(buffer.Lines) {
			buffer.Lines = append(buffer.Lines, "")
		}
		buffer.CursorX = len(buffer.Lines[len(buffer.Lines)-1])
		buffer.CursorY++
	}

	if eventKey.Key() == tcell.KeyBackspace2 {
		if err := buffer.handleBackspace(); err != nil {
			return err
		}
	}

	return nil
}

func (buffer *Buffer) handleBackspace() error {
	if buffer.CursorX <= 0 {
		if buffer.CursorY <= 0 {
			return nil
		}

		buffer.currentLine--
		buffer.CursorX = len(buffer.Lines[buffer.currentLine])
		buffer.CursorY--

		return nil
	}

	currentLine := buffer.Lines[buffer.currentLine]
	_, size := utf8.DecodeLastRune([]byte(currentLine))

	buffer.Lines[buffer.currentLine] = currentLine[:len(currentLine)-size]
	buffer.CursorX--

	return nil
}
