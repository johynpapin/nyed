package state

type State struct {
	ActiveBuffer *Buffer
	CommandLine  *CommandLine
}

func NewState() *State {
	return &State{
		ActiveBuffer: NewBuffer(),
		CommandLine:  NewCommandLine(),
	}
}
