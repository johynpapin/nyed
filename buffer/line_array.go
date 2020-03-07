package buffer

import (
	"github.com/gdamore/tcell"
)

type LineArray struct {
	lines []*Line
}

func NewLineArray() *LineArray {
	return &LineArray{}
}

func (lineArray *LineArray) ClearColors() {
	for _, line := range lineArray.lines {
		line.Colors = make(map[int]tcell.Color)
	}
}

func (lineArray *LineArray) Len() int {
	return len(lineArray.lines)
}

func (lineArray *LineArray) InsertLineBefore(lineIndex int) {
	lineArray.lines = append(lineArray.lines, nil)
	copy(lineArray.lines[lineIndex+1:], lineArray.lines[lineIndex:])
	lineArray.lines[lineIndex] = NewLine()
}

func (lineArray *LineArray) InsertLineAfter(lineIndex int) {
	lineArray.lines = append(lineArray.lines, nil)
	copy(lineArray.lines[lineIndex+2:], lineArray.lines[lineIndex+1:])
	lineArray.lines[lineIndex+1] = NewLine()
}

func (lineArray *LineArray) RemoveLine(lineIndex int) {
	lineArray.lines = append(lineArray.lines[:lineIndex], lineArray.lines[lineIndex+1:]...)

	if len(lineArray.lines) == 0 {
		lineArray.lines = []*Line{NewLine()}
	}
}

func (lineArray *LineArray) Line(lineIndex int) *Line {
	return lineArray.lines[lineIndex]
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
