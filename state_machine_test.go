package state_machine

import "testing"


func TestAddTransition(t *testing.T) {
	smachine := NewStateMachine()
	state1 := NewState("state1")
	state2 := NewState("state2")
	transition1 := NewTransition(state1, state2)

	smachine.AddTransition(*transition1)
	transitions := smachine.GetTransitions()
	if len(transitions) != 1 {
		t.Fatalf("expected 1 transition, got %d", len(transitions))
	}
}

func TestStateCanTransition(t *testing.T) {
	smachine := NewStateMachine()

	state1 := NewState("state1")
	state2 := NewState("state2")

	smachine.AddTransition(*NewTransition(state1, state2))

	_, err := smachine.Transition(state1, state2)
	if err != nil {
		t.Fatalf(
			"expected transition to be ok from %s to %s, got error: %s",
			state1.Name,
			state2.Name,
			err.Error(),
		)
	}
}

func TestStateCannotTransition(t *testing.T) {
	smachine := NewStateMachine()

	state1 := NewState("state1")
	state2 := NewState("state2")
	state3 := NewState("state3")

	smachine.AddTransition(*NewTransition(state1, state2))

	transition_ok, err := smachine.Transition(state1, state3)
	if (err == nil) && (transition_ok == true) {
		t.Fatal("transition expected to fail, got ack instead")
	}
}
