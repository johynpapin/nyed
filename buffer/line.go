package buffer

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

func (line *Line) InsertBytesAt(index int, data []byte) {
	line.data = append(line.data[:index], append(data, line.data[index:]...)...)
}

func (line *Line) InsertRuneAt(index int, r rune) {
	line.InsertBytesAt(index, []byte(string(r)))
}

func (line *Line) AppendBytes(data []byte) {
	line.data = append(line.data, data...)
}

func (line *Line) AppendRune(r rune) {
	line.AppendBytes([]byte(string(r)))
}
