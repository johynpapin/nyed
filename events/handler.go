package events

import (
	"github.com/gdamore/tcell"
	"github.com/johynpapin/nyed/highlight"
	"github.com/johynpapin/nyed/screen"
	"github.com/johynpapin/nyed/state"
	"github.com/johynpapin/nyed/ui"
)

type Handler struct {
	tcellEvents   chan tcell.Event
	waitingEvents chan Event

	Screen      *screen.Screen
	State       *state.State
	Drawer      *ui.Drawer
	Highlighter *highlight.Highlighter
}

func NewHandler(screen *screen.Screen, state *state.State, drawer *ui.Drawer, highlighter *highlight.Highlighter) *Handler {
	return &Handler{
		tcellEvents:   make(chan tcell.Event),
		waitingEvents: make(chan Event),

		Screen:      screen,
		State:       state,
		Drawer:      drawer,
		Highlighter: highlighter,
	}
}

func (handler *Handler) Start() error {
	go handler.waitingEventsLoop()
	go handler.incomingEventsLoop()
	return nil
}

func (handler *Handler) Stop() error {
	return nil
}

func (handler *Handler) PushEvent(event Event) {
	handler.waitingEvents <- event
}

func (handler *Handler) incomingEventsLoop() {
	for {
		handler.nextTcellEvent()
	}
}

func (handler *Handler) nextTcellEvent() {
	handler.Screen.Lock()
	rawEvent := handler.Screen.TcellScreen.PollEvent()
	handler.Screen.Unlock()

	switch event := rawEvent.(type) {
	case *tcell.EventKey:
		switch event.Key() {
		case tcell.KeyUp:
			handler.PushEvent(&MoveCursorEvent{
				Direction: CURSOR_DIRECTION_UP,
			})
		case tcell.KeyRight:
			handler.PushEvent(&MoveCursorEvent{
				Direction: CURSOR_DIRECTION_RIGHT,
			})
		case tcell.KeyDown:
			handler.PushEvent(&MoveCursorEvent{
				Direction: CURSOR_DIRECTION_DOWN,
			})
		case tcell.KeyLeft:
			handler.PushEvent(&MoveCursorEvent{
				Direction: CURSOR_DIRECTION_LEFT,
			})
		}
	case *tcell.EventResize:
		width, height := event.Size()
		handler.PushEvent(&ResizeEvent{
			Width:  width,
			Height: height,
		})
	}
}

func (handler *Handler) waitingEventsLoop() error {
	for {
		if err := handler.nextWaitingEvent(); err != nil {
			return err
		}
	}
}

func (handler *Handler) nextWaitingEvent() error {
	event := <-handler.waitingEvents

	if err := event.Apply(&eventContext{
		state:       handler.State,
		highlighter: handler.Highlighter,
	}); err != nil {
		return err
	}

	handler.Drawer.Draw()

	return nil
}
