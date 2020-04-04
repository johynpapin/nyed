package state

import (
	"github.com/johynpapin/nyed/utils"
	sitter "github.com/smacker/go-tree-sitter"
	"os"
)

type Buffer struct {
	Cursor    *Cursor
	LineArray *LineArray
	FilePath  string
	Tree      *sitter.Tree

	Width, Height int
	CurrentMode   utils.Mode
}

func NewBuffer() *Buffer {
	buffer := &Buffer{
		LineArray: NewLineArray(),
	}

	buffer.Cursor = NewCursor(buffer)

	return buffer
}

func (buffer *Buffer) Write() error {
	if buffer.FilePath == "" {
		return nil
	}

	file, err := os.Create(buffer.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range buffer.LineArray.Lines {
		file.Write(line.data)
		file.WriteString("\n")
	}

	return file.Sync()
}
