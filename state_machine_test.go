package state_machine

import "testing"

func TestAddAndGetState(t *testing.T) {
	smachine := NewStateMachine()

	var state1 State
	state1.Name = "Fake state 1"

	err := smachine.AddState(state1)
	if err != nil {
		t.Fatalf(err.Error())
	}

	r, _ := smachine.GetState(state1.Name)
	if r != state1 {
		t.Fatalf("expected '%s', got '%s'", state1.Name, r.Name)
	}
}

func TestAddTransition(t *testing.T) {
	smachine := NewStateMachine()
	state1 := NewState("state1")
	state2 := NewState("state2")
	transition1 := NewTransition(state1, state2)

	smachine.AddState(*state1)
	smachine.AddState(*state2)

	smachine.AddTransition(*transition1)
	transitions := smachine.GetTransitions()
	if len(transitions) != 1 {
		t.Fatalf("expected 1 transition, got %d", len(transitions))
	}
}
