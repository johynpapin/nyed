package screen

import (
	"github.com/gdamore/tcell"
	"sync"
)

type Screen struct {
	mutex *sync.Mutex

	TcellScreen tcell.Screen
}

func NewScreen() *Screen {
	return &Screen{
		mutex: &sync.Mutex{},
	}
}

func (screen *Screen) Init() error {
	var err error
	screen.TcellScreen, err = tcell.NewScreen()
	if err != nil {
		return err
	}

	return screen.TcellScreen.Init()
}

func (screen *Screen) Close() {
	screen.TcellScreen.Fini()
}

func (screen *Screen) Lock() {
	screen.mutex.Lock()
}

func (screen *Screen) Unlock() {
	screen.mutex.Unlock()
}
