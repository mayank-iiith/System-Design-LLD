package types

// Memento represents a snapshot of state. It is opaque to the caretaker.
type Memento interface {
}

// Originator defines the interface for objects that can save and restore their state.
type Originator interface {
	TakeSnapshot() Memento
	Restore(Memento)
}
