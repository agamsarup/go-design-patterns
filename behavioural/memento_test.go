package behavioural

import (
	"fmt"
	"testing"
)

func TestMemento(t *testing.T) {
	caretaker := &caretaker{
		mementoArray: make([]*memento, 0),
	}
	originator := &originator{
		state: "A",
	}
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())
	originator.setState("B")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())
	originator.setState("C")
	fmt.Printf("Originator Current State: %s\n", originator.getState())
	caretaker.addMemento(originator.createMemento())
	originator.restoreState(caretaker.getMemento(1))
	fmt.Printf("Restored to State: %s\n", originator.getState())
	originator.restoreState(caretaker.getMemento(0))
	fmt.Printf("Restored to State: %s\n", originator.getState())
}
