package ui

import (
	"github.com/gdamore/tcell"
	"github.com/johynpapin/nyed/utils"
	"github.com/mattn/go-runewidth"
	"unicode/utf8"
)

func (drawer *Drawer) drawActiveBuffer() {
	tcellScreen := drawer.Screen.TcellScreen
	buffer := drawer.State.ActiveBuffer
	lineArray := buffer.LineArray

	if buffer.Width <= 0 || buffer.Height <= 0 {
		return
	}

	currentStyle := tcell.StyleDefault

	var visualY int

LineLoop:
	for y, line := range lineArray.Lines {
		if visualY > buffer.Height {
			break LineLoop
		}

		lineBytes := line.Bytes()

		var x, visualX int
		for len(lineBytes) > 0 {
			r, size := utf8.DecodeRune(lineBytes)
			lineBytes = lineBytes[size:]

			style, exists := line.Styles[x]
			if exists {
				currentStyle = style
			}

			tcellScreen.SetContent(visualX, visualY, r, nil, currentStyle)

			if buffer.Cursor.Y == y && buffer.Cursor.X == x {
				tcellScreen.ShowCursor(visualX, visualY)
			}

			x++
			if r == '\t' {
				visualX += 8 - (visualX % 8)
			} else {
				visualX += runewidth.RuneWidth(r)
			}

			if visualX >= buffer.Width {
				visualY++
				visualX = 0

				if visualY > buffer.Height {
					break LineLoop
				}
			}
		}

		color, exists := line.Styles[x]
		if exists {
			currentStyle = color
		}

		if buffer.Cursor.Y == y && buffer.Cursor.X == x {
			tcellScreen.ShowCursor(visualX, visualY)
		}

		visualY++
	}

	if buffer.CurrentMode == utils.MODE_COMMAND {
		tcellScreen.HideCursor()
	}
}
