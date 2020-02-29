package buffer

import (
	"bufio"
	"github.com/gdamore/tcell"
	"github.com/johynpapin/nyed/screen"
	"github.com/johynpapin/nyed/statusline"
	"github.com/johynpapin/nyed/ui"
	"github.com/johynpapin/nyed/utils"
	"github.com/mattn/go-runewidth"
	"os"
	"unicode/utf8"
)

type Buffer struct {
	ui.Section

	statusLine     *statusline.StatusLine
	lineArray      *LineArray
	Cursor         *Cursor
	currentCommand string

	FilePath string
}

func NewBuffer() *Buffer {
	buffer := &Buffer{
		statusLine: statusline.NewStatusLine(),
		lineArray:  NewLineArray(),
	}
	buffer.Cursor = NewCursor(buffer)

	return buffer
}

func (buffer *Buffer) Draw(screen *screen.Screen) error {
	if buffer.Width <= 0 || buffer.Height <= 0 {
		return nil
	}

	var visualY int

LineLoop:
	for y, line := range buffer.lineArray.lines {
		if visualY > buffer.Height {
			break LineLoop
		}

		lineData := line.data

		var x, visualX int
		for len(lineData) > 0 {
			r, size := utf8.DecodeRune(lineData)
			lineData = lineData[size:]

			screen.Screen.SetContent(visualX, visualY, r, nil, tcell.StyleDefault)

			if buffer.Cursor.Y == y && buffer.Cursor.X == x {
				screen.Screen.ShowCursor(visualX, visualY)
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

		if buffer.Cursor.Y == y && buffer.Cursor.X == x {
			screen.Screen.ShowCursor(visualX, visualY)
		}

		visualY++
	}

	if buffer.CurrentMode() == utils.MODE_COMMAND {
		screen.Screen.HideCursor()
	}

	return buffer.statusLine.Draw(screen)
}

func (buffer *Buffer) HandleEventKey(eventKey *tcell.EventKey) error {
	switch buffer.CurrentMode() {
	case utils.MODE_INSERT:
		if eventKey.Key() == tcell.KeyRune {
			buffer.lineArray.Line(buffer.Cursor.Y).InsertRuneAt(buffer.Cursor.X, eventKey.Rune())
			buffer.Cursor.MoveRight()
		}

		if eventKey.Key() == tcell.KeyTab {
			buffer.lineArray.Line(buffer.Cursor.Y).InsertRuneAt(buffer.Cursor.X, '\t')
			buffer.Cursor.MoveRight()
		}

		if eventKey.Key() == tcell.KeyEnter {
			buffer.lineArray.SplitLineAt(buffer.Cursor.Y, buffer.Cursor.X)
			buffer.Cursor.X = 0
			buffer.Cursor.savedVisualX = 0
			buffer.Cursor.Y++
		}

		if eventKey.Key() == tcell.KeyBackspace2 {
			if err := buffer.handleBackspace(); err != nil {
				return err
			}
		}

		if eventKey.Key() == tcell.KeyEscape {
			buffer.SetCurrentMode(utils.MODE_NORMAL)
			buffer.Cursor.MoveLeft()
		}

	case utils.MODE_NORMAL:
		if eventKey.Key() == tcell.KeyRune {
			switch eventKey.Rune() {
			case 'a':
				buffer.SetCurrentMode(utils.MODE_INSERT)
				buffer.Cursor.MoveRight()
			case 'i':
				buffer.SetCurrentMode(utils.MODE_INSERT)
			case ':':
				buffer.SetCurrentMode(utils.MODE_COMMAND)
			case '$':
				buffer.Cursor.MoveToEndOfLine()
			case 'o':
				buffer.lineArray.InsertLineAfter(buffer.Cursor.Y)
				buffer.Cursor.Y++
				buffer.Cursor.X = 0
				buffer.Cursor.savedVisualX = 0
				buffer.SetCurrentMode(utils.MODE_INSERT)
			case 'O':
				buffer.lineArray.InsertLineBefore(buffer.Cursor.Y)
				buffer.Cursor.X = 0
				buffer.Cursor.savedVisualX = 0
				buffer.SetCurrentMode(utils.MODE_INSERT)
			case 'd':
				if buffer.currentCommand == "d" {
					buffer.lineArray.RemoveLine(buffer.Cursor.Y)
					buffer.currentCommand = ""

					if buffer.Cursor.Y >= len(buffer.lineArray.lines) {
						buffer.Cursor.MoveUp()
					}

					buffer.Cursor.clamp()
					utils.ResetCursorStyle()
					break
				}

				buffer.currentCommand = "d"
				utils.SetCursorStyle(utils.CURSOR_STYLE_BLINKING_UNDERLINE)
			}
		}
	}

	switch eventKey.Key() {
	case tcell.KeyEnd:
		buffer.Cursor.MoveToEndOfLine()
	case tcell.KeyLeft:
		buffer.Cursor.MoveLeft()
	case tcell.KeyRight:
		buffer.Cursor.MoveRight()
	case tcell.KeyUp:
		buffer.Cursor.MoveUp()
	case tcell.KeyDown:
		buffer.Cursor.MoveDown()
	}

	return nil
}

func (buffer *Buffer) handleBackspace() error {
	if buffer.Cursor.X <= 0 {
		if buffer.Cursor.Y <= 0 {
			return nil
		}

		buffer.Cursor.Y--
		buffer.Cursor.X = len(buffer.lineArray.lines[buffer.Cursor.Y].data)

		buffer.lineArray.MergeLineAtTheEndOf(buffer.Cursor.Y, buffer.Cursor.Y+1)

		return nil
	}

	buffer.lineArray.Line(buffer.Cursor.Y).RemoveRune(buffer.Cursor.X - 1)
	buffer.Cursor.MoveLeft()

	return nil
}

func (buffer *Buffer) SetSize(x, y, width, height int) {
	buffer.X = x
	buffer.Y = y
	buffer.Width = width
	buffer.Height = height - 1

	buffer.statusLine.X = x
	buffer.statusLine.Y = height - 1
	buffer.statusLine.Width = width
	buffer.statusLine.Height = 1
}

func (buffer *Buffer) Load() error {
	if buffer.FilePath == "" {
		buffer.lineArray.lines = []*Line{NewLine()}
		return nil
	}

	file, err := os.Open(buffer.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buffer.lineArray.lines = append(buffer.lineArray.lines, NewLineFromBytes(scanner.Bytes()))
	}

	return scanner.Err()
}

func (buffer *Buffer) Write() error {
	if buffer.FilePath == "" {
		return nil
	}

	file, err := os.Create(buffer.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range buffer.lineArray.lines {
		file.Write(line.data)
		file.WriteString("\n")
	}

	return file.Sync()
}

func (buffer *Buffer) SetCurrentMode(mode utils.Mode) {
	buffer.statusLine.CurrentMode = mode

	switch mode {
	case utils.MODE_INSERT:
		utils.SetCursorStyle(utils.CURSOR_STYLE_BLINKING_BAR)
	case utils.MODE_NORMAL:
		fallthrough
	case utils.MODE_COMMAND:
		utils.ResetCursorStyle()
	}
}

func (buffer *Buffer) CurrentMode() utils.Mode {
	return buffer.statusLine.CurrentMode
}
