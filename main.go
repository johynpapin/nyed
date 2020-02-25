package main

import (
	"fmt"
	"github.com/johynpapin/nyed/buffer"
	"github.com/johynpapin/nyed/events"
	"github.com/johynpapin/nyed/statusline"
	"github.com/johynpapin/nyed/ui"
	"os"
)

func checkError(err error) {
	if err == nil {
		return
	}

	fmt.Fprintf(os.Stderr, "%v\n", err)
	os.Exit(1)
}

func main() {
	screen := ui.NewScreen()
	handler := events.NewHandler(screen)

	checkError(screen.Init())
	checkError(handler.Start())

	handler.CurrentBuffer = buffer.NewBuffer()
	screen.AddDrawable(handler.CurrentBuffer)

	screen.AddDrawable(statusline.NewStatusLine())

	for {
		checkError(screen.Draw())
		checkError(handler.Next())
	}
}
