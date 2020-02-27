package utils

import (
	"fmt"
)

type CursorStyle int

const (
	CURSOR_STYLE_BLINKING_BLOCK CursorStyle = iota
	CURSOR_STYLE_DEFAULT
	CURSOR_STYLE_STEADY_BLOCK
	CURSOR_STYLE_BLINKING_UNDERLINE
	CURSOR_STYLE_STEADY_UNDERLINE
	CURSOR_STYLE_BLINKING_BAR
	CURSOR_STYLE_STEADY_BAR
)

func SetCursorStyle(cursorStyle CursorStyle) {
	fmt.Printf("\033[%d q", cursorStyle)
}

func ResetCursorStyle() {
	SetCursorStyle(CURSOR_STYLE_DEFAULT)
}
