package statemachine

import (
	"fmt"
)

// StateMachine contains all the transitions available
type StateMachine struct {
	transitions []Transition
}

// NewStateMachine returns a new instance of the StateMachine struct
func NewStateMachine() *StateMachine {
	return &StateMachine{}
}

// AddTransition adds a new transition to the state machine
func (state_machine *StateMachine) AddTransition(transition Transition) {
	state_machine.transitions = append(state_machine.transitions, transition)
}

// GetTransitions retrievs all the available transitions on the state machine
func (state_machine *StateMachine) GetTransitions() []Transition {
	return state_machine.transitions
}

// Transition returns if you can transition between two states
func (state_machine *StateMachine) Transition(from *State, to *State) (bool, error) {
	transitions := state_machine.GetTransitions()
	for _, transition := range transitions {
		if (transition.From.Name == from.Name) && (transition.To.Name == to.Name) {
			return true, nil
		}
	}

	return false, fmt.Errorf("Cannot transition from %s to %s", from.Name, to.Name)
}
