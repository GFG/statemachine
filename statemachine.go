package statemachine

import (
	"errors"
	"fmt"
)

type StateMachine struct {
	transitions []Transition
}

func NewStateMachine() *StateMachine {
	return &StateMachine{}
}

func (state_machine *StateMachine) AddTransition(transition Transition) {
	state_machine.transitions = append(state_machine.transitions, transition)
}

func (state_machine *StateMachine) GetTransitions() []Transition {
	return state_machine.transitions
}

func (state_machine *StateMachine) Transition(from *State, to *State) (bool, error) {
	transitions := state_machine.GetTransitions()
	for _, transition := range transitions {
		if (transition.From.Name == from.Name) && (transition.To.Name == to.Name) {
			return true, nil
		}
	}

	return false, errors.New(fmt.Sprintf("Cannot transition from %s to %s", from.Name, to.Name))
}
