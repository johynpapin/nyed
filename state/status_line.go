package state

type StatusLine struct {
	Buffer *Buffer
}

func newStatusLine() *StatusLine {
	return &StatusLine{}
}

func (statusLine *StatusLine) CurrentMode() {

}
