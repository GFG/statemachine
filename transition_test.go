package state_machine

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewTransition(t *testing.T) {
	state1 := NewState("State1")
	state2 := NewState("State2")

	transition := NewTransition(state1, state2)

	assert.Equal(t, transition.From, state1)
	assert.Equal(t, transition.To, state2)
}
