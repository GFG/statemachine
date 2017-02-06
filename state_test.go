package state_machine

import (
    "testing"
)

func TestNewState(t *testing.T) {
    state_name := "Fake test"
    state := NewState(state_name)

    if state.Name != state_name {
        t.Fatalf("expected '%s', got '%s'", state_name, state.Name)
    }
}