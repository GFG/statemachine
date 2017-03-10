package statemachine

// Transition contains a two state flow
type Transition struct {
	From *State
	To   *State
}

// NewTransition returns a new instance of the Transition structure
func NewTransition(from *State, to *State) *Transition {
	return &Transition{
		From: from,
		To:   to,
	}
}
