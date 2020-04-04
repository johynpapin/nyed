package main

import (
	"github.com/johynpapin/nyed/events"
	"github.com/johynpapin/nyed/highlight"
	"github.com/johynpapin/nyed/screen"
	"github.com/johynpapin/nyed/state"
	"github.com/johynpapin/nyed/ui"
)

type Editor struct {
	state         *state.State
	screen        *screen.Screen
	eventsHandler *events.Handler
	drawer        *ui.Drawer
	highlighter   *highlight.Highlighter
}

func NewEditor(files []string) *Editor {
	editor := &Editor{
		state:  state.NewState(),
		screen: screen.NewScreen(),
	}

	editor.drawer = ui.NewDrawer(editor.screen, editor.state)
	editor.highlighter = highlight.NewHighlighter(editor.state.ActiveBuffer)

	if len(files) >= 1 {
		editor.state.ActiveBuffer.FilePath = files[0]
	}

	editor.eventsHandler = events.NewHandler(editor.screen, editor.state, editor.drawer, editor.highlighter)

	return editor
}

func (editor *Editor) Start() error {
	if err := editor.screen.Init(); err != nil {
		return err
	}

	editor.highlighter.Init()

	if err := editor.eventsHandler.Start(); err != nil {
		return err
	}

	editor.eventsHandler.PushEvent(&events.LoadEvent{})

	<-make(chan struct{})

	return nil
}

func (editor *Editor) Stop() error {
	if err := editor.eventsHandler.Stop(); err != nil {
		return err
	}

	editor.screen.Close()

	return nil
}
