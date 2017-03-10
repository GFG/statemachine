package statemachine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewState(t *testing.T) {
	state_name := "Fake test"
	state := NewState(state_name)

	assert.Equal(t, state_name, state.Name)
}
