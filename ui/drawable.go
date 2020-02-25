package ui

type Drawable interface {
	Draw(*Screen) error
}
