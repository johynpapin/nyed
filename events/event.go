package events

import (
	"github.com/johynpapin/nyed/highlight"
	"github.com/johynpapin/nyed/state"
)

type eventContext struct {
	state       *state.State
	highlighter *highlight.Highlighter
}

type Event interface {
	Apply(*eventContext) error
}
