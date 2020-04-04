package events

import (
	"bufio"
	"github.com/johynpapin/nyed/state"
	"os"
)

type LoadEvent struct{}

func (e *LoadEvent) Apply(ctx *eventContext) error {
	buffer := ctx.state.ActiveBuffer

	if buffer.FilePath == "" {
		buffer.LineArray.Lines = []*state.Line{state.NewLine()}
		return nil
	}

	file, err := os.Open(buffer.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		buffer.LineArray.Lines = append(buffer.LineArray.Lines, state.NewLineFromBytes(scanner.Bytes()))
	}

	if err = scanner.Err(); err != nil {
		return err
	}

	ctx.highlighter.Buffer = buffer
	return ctx.highlighter.ReHighlight()
}
