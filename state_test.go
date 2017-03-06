package statemachine

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewState(t *testing.T) {
	state_name := "Fake test"
	state := NewState(state_name)

	assert.Equal(t, state_name, state.Name)
}
