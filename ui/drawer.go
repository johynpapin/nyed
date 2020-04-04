package ui

import (
	"github.com/johynpapin/nyed/screen"
	"github.com/johynpapin/nyed/state"
)

type Drawer struct {
	Screen *screen.Screen
	State  *state.State
}

func NewDrawer(screen *screen.Screen, state *state.State) *Drawer {
	return &Drawer{
		Screen: screen,
		State:  state,
	}
}

func (drawer *Drawer) Draw() {
	tcellScreen := drawer.Screen.TcellScreen

	tcellScreen.Clear()
	tcellScreen.HideCursor()

	drawer.drawActiveBuffer()
	drawer.drawCommandLine()

	tcellScreen.Show()
}
