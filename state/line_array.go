package state

import (
	"github.com/gdamore/tcell"
)

type LineArray struct {
	Lines []*Line
}

func NewLineArray() *LineArray {
	return &LineArray{}
}

func (lineArray *LineArray) ClearStyles() {
	for _, line := range lineArray.Lines {
		line.Styles = make(map[int]tcell.Style)
	}
}

func (lineArray *LineArray) Len() int {
	return len(lineArray.Lines)
}

func (lineArray *LineArray) InsertLineBefore(lineIndex int) {
	lineArray.Lines = append(lineArray.Lines, nil)
	copy(lineArray.Lines[lineIndex+1:], lineArray.Lines[lineIndex:])
	lineArray.Lines[lineIndex] = NewLine()
}

func (lineArray *LineArray) InsertLineAfter(lineIndex int) {
	lineArray.Lines = append(lineArray.Lines, nil)
	copy(lineArray.Lines[lineIndex+2:], lineArray.Lines[lineIndex+1:])
	lineArray.Lines[lineIndex+1] = NewLine()
}

func (lineArray *LineArray) RemoveLine(lineIndex int) {
	lineArray.Lines = append(lineArray.Lines[:lineIndex], lineArray.Lines[lineIndex+1:]...)

	if len(lineArray.Lines) == 0 {
		lineArray.Lines = []*Line{NewLine()}
	}
}

func (lineArray *LineArray) Line(lineIndex int) *Line {
	return lineArray.Lines[lineIndex]
}

func (lineArray *LineArray) MergeLineAtTheEndOf(sourceLineIndex int, targetLineIndex int) {
	sourceLine := lineArray.Line(sourceLineIndex)
	targetLine := lineArray.Line(targetLineIndex)

	sourceLine.AppendBytes(targetLine.data)

	lineArray.RemoveLine(targetLineIndex)
}

func (lineArray *LineArray) SplitLineAt(lineIndex int, x int) {
	lineToSplit := lineArray.Line(lineIndex)
	lineArray.InsertLineAfter(lineIndex)
	targetLine := lineArray.Line(lineIndex + 1)

	if x <= 0 {
		targetLine.data = lineToSplit.data
		lineToSplit.data = nil
		return
	}

	byteIndex := lineToSplit.findByteIndexFromRuneIndex(x - 1)
	targetLine.data = lineToSplit.data[byteIndex+1:]
	lineToSplit.data = append([]byte(nil), lineToSplit.data[:byteIndex+1]...)
}
