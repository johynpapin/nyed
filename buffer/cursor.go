package buffer

import (
	"github.com/johynpapin/nyed/utils"
)

type Cursor struct {
	buffer *Buffer
	X, Y   int
	savedX int
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
	cursor.savedX = cursor.X
}

func (cursor *Cursor) MoveRight() {
	cursorLimitX := cursor.cursorLimitX()

	if cursor.X >= cursorLimitX {
		return
	}

	cursor.X++
	cursor.savedX = cursor.X
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

func (cursor *Cursor) MoveToEndOfLine() {
	cursor.X = cursor.cursorLimitX()
}

func (cursor *Cursor) clamp() {
	cursorLimitX := cursor.cursorLimitX()

	if cursor.savedX >= cursorLimitX {
		cursor.X = cursorLimitX
	}
}

func (cursor *Cursor) cursorLimitX() int {
	cursorLimitX := len(string(cursor.buffer.lineArray.Line(cursor.Y).data)) - 1

	if cursor.buffer.CurrentMode() == utils.MODE_INSERT {
		cursorLimitX++
	}

	if cursorLimitX < 0 {
		return 0
	}

	return cursorLimitX
}
