package statemachine

type State struct {
	Name string
}

func NewState(name string) *State {
	return &State{
		Name: name,
	}
}
