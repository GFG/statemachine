package statemachine

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewState(t *testing.T) {
	stateName := "Fake test"
	state := NewState(stateName)

	assert.Equal(t, stateName, state.Name)
}
