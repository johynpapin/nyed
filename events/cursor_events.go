package events

type CursorDirection int

const (
	CURSOR_DIRECTION_UP CursorDirection = iota
	CURSOR_DIRECTION_RIGHT
	CURSOR_DIRECTION_DOWN
	CURSOR_DIRECTION_LEFT
)

type MoveCursorEvent struct {
	Direction CursorDirection
}

func (e *MoveCursorEvent) Apply(ctx *eventContext) error {
	cursor := ctx.state.ActiveBuffer.Cursor

	switch e.Direction {
	case CURSOR_DIRECTION_UP:
		cursor.MoveUp()
	case CURSOR_DIRECTION_RIGHT:
		cursor.MoveRight()
	case CURSOR_DIRECTION_DOWN:
		cursor.MoveDown()
	case CURSOR_DIRECTION_LEFT:
		cursor.MoveLeft()
	}

	return nil
}
