package buffer

import (
	"github.com/johynpapin/nyed/utils"
	"github.com/mattn/go-runewidth"
	"unicode/utf8"
)

type Cursor struct {
	buffer       *Buffer
	X, Y         int
	savedVisualX int
}

func NewCursor(buffer *Buffer) *Cursor {
	return &Cursor{
		buffer: buffer,
	}
}

func (cursor *Cursor) MoveLeft() {
	if cursor.X <= 0 {
		return
	}

	cursor.X--
	cursor.saveVisualX()
}

func (cursor *Cursor) MoveRight() {
	cursorLimitX := cursor.cursorLimitX()

	if cursor.X >= cursorLimitX {
		return
	}

	cursor.X++
	cursor.saveVisualX()
}

func (cursor *Cursor) MoveUp() {
	if cursor.Y <= 0 {
		return
	}

	cursor.Y--
	cursor.clamp()
}

func (cursor *Cursor) MoveDown() {
	if cursor.Y >= len(cursor.buffer.lineArray.lines)-1 {
		return
	}

	cursor.Y++
	cursor.clamp()
}

func (cursor *Cursor) MoveToStartOfLine() {
	cursor.X = 0
	cursor.savedVisualX = 0
}

func (cursor *Cursor) MoveToEndOfLine() {
	cursor.X = cursor.cursorLimitX()
}

func (cursor *Cursor) clamp() {
	cursor.X = cursor.xFromVisualX(cursor.savedVisualX, 8)

	cursorLimitX := cursor.cursorLimitX()

	if cursor.X > cursorLimitX {
		cursor.X = cursorLimitX
	}
}

func (cursor *Cursor) cursorLimitX() int {
	cursorLimitX := cursor.buffer.lineArray.Line(cursor.Y).LengthInRunes() - 1

	if cursor.buffer.CurrentMode() == utils.MODE_INSERT {
		cursorLimitX++
	}

	if cursorLimitX < 0 {
		return 0
	}

	return cursorLimitX
}

func (cursor *Cursor) saveVisualX() {
	cursor.savedVisualX = cursor.visualX()
}

func (cursor *Cursor) visualX() int {
	if cursor.X <= 0 {
		return 0
	}

	return cursor.buffer.lineArray.Line(cursor.Y).VisualWidth(cursor.X-1, 8)
}

func (cursor *Cursor) xFromVisualX(visualX int, tabWidth int) int {
	data := cursor.buffer.lineArray.Line(cursor.Y).data

	var x, currentVisualX int
	for len(data) > 0 && visualX > currentVisualX {
		r, size := utf8.DecodeRune(data)
		data = data[size:]

		if r == '\t' {
			currentVisualX += tabWidth - (currentVisualX % tabWidth)
		} else {
			currentVisualX += runewidth.RuneWidth(r)
		}

		x++
	}

	if currentVisualX > visualX {
		x--
	}

	return x
}
