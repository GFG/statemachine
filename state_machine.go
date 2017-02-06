package state_machine

import "errors"

type StateMachine struct {
	states      map[string]State
	transitions []Transition
}

func NewStateMachine() *StateMachine {
	return &StateMachine{
		states: make(map[string]State),
	}
}

func (state_machine *StateMachine) GetState(name string) (State, error) {
	var state State

	state, err := state_machine.states[name]
	if !err {
		return state, errors.New("State not found")
	}

	return state, nil
}

func (state_machine *StateMachine) AddState(state State) error {
	_, err := state_machine.GetState(state.Name)
	if err == nil {
		return errors.New("State already exists")
	}

	state_machine.states[state.Name] = state

	return nil
}

func (state_machine *StateMachine) AddTransition(transition Transition) {
	state_machine.transitions = append(state_machine.transitions, transition)
}

func (state_machine *StateMachine) GetTransitions() []Transition {
	return state_machine.transitions
}
