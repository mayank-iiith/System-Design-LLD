package caretaker

import (
	"lld/behavioral_design_pattens/memento/types"
)

// History acts as caretaker. It stores undo/redo stacks.
type History struct {
	undos []types.Memento
	redos []types.Memento
}

// Save current state of the originator
func (h *History) Save(o types.Originator) {
	h.undos = append(h.undos, o.TakeSnapshot())
	h.redos = nil // clear redos when a new action occurs
}

// Undo reverts to the previous state
func (h *History) Undo(o types.Originator) bool {
	if len(h.undos) == 0 {
		return false
	}
	last := h.undos[len(h.undos)-1]
	h.undos = h.undos[:len(h.undos)-1]
	h.redos = append(h.redos, o.TakeSnapshot())
	o.Restore(last)
	return true
}

// Redo reapplies a reverted state
func (h *History) Redo(o types.Originator) bool {
	if len(h.redos) == 0 {
		return false
	}
	last := h.redos[len(h.redos)-1]
	h.redos = h.redos[:len(h.redos)-1]
	h.undos = append(h.undos, o.TakeSnapshot())
	o.Restore(last)
	return true
}
