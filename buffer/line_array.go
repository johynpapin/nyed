package buffer

type LineArray struct {
	lines []*Line
}

func NewLineArray() *LineArray {
	return &LineArray{}
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
