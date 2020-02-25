package ui

import (
	"github.com/gdamore/tcell"
	"sync"
)

type Screen struct {
	mutex *sync.Mutex

	Screen tcell.Screen

	drawables []Drawable
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

func (screen *Screen) Draw() error {
	screen.Screen.Clear()
	screen.Screen.HideCursor()

	for _, drawable := range screen.drawables {
		if err := drawable.Draw(screen); err != nil {
			return err
		}
	}

	screen.Screen.Show()

	return nil
}

func (screen *Screen) AddDrawable(drawable Drawable) {
	screen.drawables = append(screen.drawables, drawable)
}

func (screen *Screen) Lock() {
	screen.mutex.Lock()
}

func (screen *Screen) Unlock() {
	screen.mutex.Unlock()
}
