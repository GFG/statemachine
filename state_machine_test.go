package state_machine

import (
    "testing"
)

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