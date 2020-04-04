package events

import (
	"github.com/johynpapin/nyed/state"
)

type InsertEvent struct {
	Buffer *state.Buffer
}

func (e *InsertEvent) Apply(ctx *eventContext) error {
	return nil
}
