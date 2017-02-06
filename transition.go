package state_machine

type Transition struct {
    From *State
    To   *State
}

func NewTransition(from *State, to *State) *Transition {
    return &Transition{
        From: from,
        To: to,
    }
}