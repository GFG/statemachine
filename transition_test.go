package state_machine

import "testing"

func TestNewTransition(t *testing.T) {
	state1 := NewState("State1")
	state2 := NewState("State2")

	transition := NewTransition(state1, state2)

	if transition.From != state1 {
		t.Fatalf("expected '%s', got '%s'", state1.Name, transition.From.Name)
	}

	if transition.To != state2 {
		t.Fatalf("expected '%s', got '%s'", state2.Name, transition.To.Name)
	}
}
