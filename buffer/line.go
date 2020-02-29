package buffer

import (
	"github.com/mattn/go-runewidth"
	"unicode/utf8"
)

type Line struct {
	data []byte
}

func NewLine() *Line {
	return &Line{}
}

func NewLineFromBytes(data []byte) *Line {
	return &Line{
		data: data,
	}
}

func (line *Line) findByteIndexFromRuneIndex(runeIndex int) int {
	if runeIndex == 0 {
		return 0
	}

	byteIndex := 0
	data := line.data
	for runeIndex > 0 && len(data) > 0 {
		_, size := utf8.DecodeRune(data)
		data = data[size:]

		byteIndex += size
		runeIndex--
	}

	return byteIndex
}

func (line *Line) LengthInRunes() int {
	return utf8.RuneCount(line.data)
}

func (line *Line) InsertBytesAt(runeIndex int, data []byte) {
	index := line.findByteIndexFromRuneIndex(runeIndex)
	line.data = append(line.data[:index], append(data, line.data[index:]...)...)
}

func (line *Line) InsertRuneAt(runeIndex int, r rune) {
	line.InsertBytesAt(runeIndex, []byte(string(r)))
}

func (line *Line) AppendBytes(data []byte) {
	line.data = append(line.data, data...)
}

func (line *Line) AppendRune(r rune) {
	line.AppendBytes([]byte(string(r)))
}

func (line *Line) VisualWidth(runeIndex int, tabWidth int) int {
	data := line.data
	visualWidth := 0
	for len(data) > 0 && runeIndex >= 0 {
		r, size := utf8.DecodeRune(data)
		data = data[size:]

		if r == '\t' {
			visualWidth += tabWidth - (visualWidth % tabWidth)
		} else {
			visualWidth += runewidth.RuneWidth(r)
		}

		runeIndex--
	}

	return visualWidth
}

func (line *Line) RemoveRune(runeIndex int) {
	line.data = append(line.data[:runeIndex], line.data[runeIndex+1:]...)
}
