package events

import (
	"github.com/gdamore/tcell"
)

func (handler *Handler) handleResize(event *tcell.EventResize) {
	width, height := event.Size()

	handler.CurrentBuffer.SetSize(0, 0, width, height-1)

	handler.CommandLine.X = 0
	handler.CommandLine.Y = height - 1
	handler.CommandLine.Width = width
	handler.CommandLine.Height = 1
}
