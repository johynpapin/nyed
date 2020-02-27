package screen

import (
	"github.com/gdamore/tcell"
	"sync"
)

type Screen struct {
	mutex *sync.Mutex

	Screen tcell.Screen
}

func NewScreen() *Screen {
	return &Screen{
		mutex: &sync.Mutex{},
	}
}

func (screen *Screen) Init() error {
	var err error
	screen.Screen, err = tcell.NewScreen()
	if err != nil {
		return err
	}

	return screen.Screen.Init()
}

func (screen *Screen) Close() {
	screen.Screen.Fini()
}

func (screen *Screen) Lock() {
	screen.mutex.Lock()
}

func (screen *Screen) Unlock() {
	screen.mutex.Unlock()
}
