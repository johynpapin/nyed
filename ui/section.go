package ui

type Section struct {
	X, Y          int
	Width, Height int
	Drawable      Drawable
}

func NewSection() *Section {
	return &Section{}
}
