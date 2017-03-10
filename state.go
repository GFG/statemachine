package statemachine

// State contains the name of the state
type State struct {
	Name string
}

// NewState returns a new instance of the State struct
func NewState(name string) *State {
	return &State{
		Name: name,
	}
}
