package statemachine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTransition(t *testing.T) {
	smachine := NewStateMachine()
	state1 := NewState("state1")
	state2 := NewState("state2")
	transition1 := NewTransition(state1, state2)

	smachine.AddTransition(*transition1)
	transitions := smachine.GetTransitions()

	assert.Equal(t, 1, len(transitions))
}

func TestStateCanTransition(t *testing.T) {
	smachine := NewStateMachine()

	state1 := NewState("state1")
	state2 := NewState("state2")

	smachine.AddTransition(*NewTransition(state1, state2))

	ack, err := smachine.Transition(state1, state2)
	assert.Equal(t, true, ack)
	assert.Nil(t, err)
}

func TestStateCannotTransition(t *testing.T) {
	smachine := NewStateMachine()

	state1 := NewState("state1")
	state2 := NewState("state2")
	state3 := NewState("state3")

	smachine.AddTransition(*NewTransition(state1, state2))

	transition_ok, err := smachine.Transition(state1, state3)
	assert.Equal(t, false, transition_ok)
	assert.NotNil(t, err)
}
