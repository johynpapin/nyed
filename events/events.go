package events

import (
	"github.com/gdamore/tcell"
	"github.com/johynpapin/nyed/buffer"
	"github.com/johynpapin/nyed/ui"
	"os"
)

type Handler struct {
	screen *ui.Screen

	events chan tcell.Event

	CurrentBuffer *buffer.Buffer
}

func NewHandler(screen *ui.Screen) *Handler {
	return &Handler{
		screen: screen,

		events: make(chan tcell.Event),
	}
}

func (handler *Handler) Start() error {
	go func() {
		for {
			handler.screen.Lock()
			event := handler.screen.Screen.PollEvent()
			handler.screen.Unlock()

			handler.events <- event
		}
	}()

	return nil
}

func (handler *Handler) Stop() error {
	return nil
}

func (handler *Handler) Next() error {
	event := <-handler.events

	if eventKey, ok := event.(*tcell.EventKey); ok {
		if eventKey.Key() == tcell.KeyCtrlC || eventKey.Key() == tcell.KeyCtrlQ {
			handler.screen.Screen.Fini()
			os.Exit(0)
		}

		if err := handler.CurrentBuffer.HandleEventKey(eventKey); err != nil {
			return err
		}
	}

	return nil
}
