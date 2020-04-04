package events

type ResizeEvent struct {
	Width, Height int
}

func (e *ResizeEvent) Apply(ctx *eventContext) error {
	state := ctx.state

	state.ActiveBuffer.Height = e.Height - 1
	state.ActiveBuffer.Width = e.Width

	state.CommandLine.Y = e.Height - 1
	state.CommandLine.Width = e.Width

	return nil
}
