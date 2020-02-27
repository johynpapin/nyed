package main

import (
	"github.com/johynpapin/nyed/buffer"
	"github.com/johynpapin/nyed/commandline"
	"github.com/johynpapin/nyed/events"
	"github.com/johynpapin/nyed/screen"
	"os"
)

type Editor struct {
	screen       *screen.Screen
	eventHandler *events.Handler

	currentBuffer *buffer.Buffer
	commandLine   *commandline.CommandLine
}

func NewEditor(files []string) *Editor {
	editor := &Editor{
		screen: screen.NewScreen(),

		currentBuffer: buffer.NewBuffer(),
		commandLine:   commandline.NewCommandLine(),
	}

	if len(files) >= 1 {
		editor.currentBuffer.FilePath = files[0]
	}

	editor.eventHandler = events.NewHandler(editor.screen, editor.commandLine, func() {
		editor.Stop()
		os.Exit(0)
	})
	editor.eventHandler.CurrentBuffer = editor.currentBuffer

	return editor
}

func (editor *Editor) Start() error {
	if err := editor.currentBuffer.Load(); err != nil {
		return err
	}

	if err := editor.screen.Init(); err != nil {
		return err
	}

	if err := editor.eventHandler.Start(); err != nil {
		return err
	}

	return editor.mainLoop()
}

func (editor *Editor) Stop() error {
	if err := editor.eventHandler.Stop(); err != nil {
		return err
	}

	editor.screen.Close()

	return nil
}

func (editor *Editor) mainLoop() error {
	for {
		if err := editor.eventHandler.Next(); err != nil {
			return err
		}

		if err := editor.draw(); err != nil {
			return err
		}
	}
}

func (editor *Editor) draw() error {
	screen := editor.screen.Screen

	screen.Clear()
	screen.HideCursor()

	if err := editor.currentBuffer.Draw(editor.screen); err != nil {
		return err
	}

	if err := editor.commandLine.Draw(editor.screen); err != nil {
		return err
	}

	screen.Show()

	return nil
}
