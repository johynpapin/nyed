package events

import (
	"github.com/gdamore/tcell"
	"github.com/johynpapin/nyed/buffer"
	"github.com/johynpapin/nyed/commandline"
	"github.com/johynpapin/nyed/screen"
	"github.com/johynpapin/nyed/utils"
)

type Handler struct {
	events chan tcell.Event

	Screen        *screen.Screen
	CommandLine   *commandline.CommandLine
	onQuit        func()
	CurrentBuffer *buffer.Buffer
}

func NewHandler(screen *screen.Screen, commandLine *commandline.CommandLine, onQuit func()) *Handler {
	return &Handler{
		events: make(chan tcell.Event),

		Screen:      screen,
		CommandLine: commandLine,
		onQuit:      onQuit,
	}
}

func (handler *Handler) Start() error {
	go func() {
		for {
			handler.Screen.Lock()
			event := handler.Screen.Screen.PollEvent()
			handler.Screen.Unlock()

			handler.events <- event
		}
	}()

	return nil
}

func (handler *Handler) Stop() error {
	return nil
}

func (handler *Handler) Next() error {
	rawEvent := <-handler.events

	switch event := rawEvent.(type) {
	case *tcell.EventKey:
		if handler.CurrentBuffer.CurrentMode() != utils.MODE_COMMAND {
			if err := handler.CurrentBuffer.HandleEventKey(event); err != nil {
				return err
			}

			if handler.CurrentBuffer.CurrentMode() == utils.MODE_COMMAND {
				handler.CommandLine.CurrentCommand = ":"
			}
		} else {
			command, err := handler.CommandLine.HandleEventKey(event)
			if err != nil {
				return err
			}

			if command != "" {
				handler.handleCommand(command)
			}

			if handler.CommandLine.CurrentCommand == "" {
				handler.CurrentBuffer.SetCurrentMode(utils.MODE_NORMAL)
				utils.ResetCursorStyle()
			}
		}
	case *tcell.EventResize:
		handler.handleResize(event)
	}

	return nil
}

func (handler *Handler) handleCommand(command string) {
	switch command {
	case "w":
		handler.CurrentBuffer.Write()
	case "q":
		fallthrough
	case "q!":
		fallthrough
	case "x":
		handler.onQuit()
	}
}
