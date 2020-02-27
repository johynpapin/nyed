package utils

import (
	"github.com/gdamore/tcell"
)

type Mode int

const (
	MODE_NORMAL Mode = iota
	MODE_INSERT
	MODE_VISUAL
	MODE_COMMAND
)

func GetModeStyle(mode Mode) tcell.Style {
	return tcell.StyleDefault.Background(tcell.ColorLightBlue).Foreground(tcell.ColorBlack).Bold(true)
}

func GetModeText(mode Mode) string {
	switch mode {
	case MODE_NORMAL:
		return "NORMAL"
	case MODE_INSERT:
		return "INSERT"
	case MODE_VISUAL:
		return "VISUAL"
	case MODE_COMMAND:
		return "COMMAND"
	}

	return ""
}
