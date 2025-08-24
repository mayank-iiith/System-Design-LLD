package history

import (
	"errors"
	"lld/behavioral_design_pattens/command/commands"
)

// Note: We can use Memento Design Pattern here for history

type CommandHistory struct {
	history []commands.Command
}

func NewCommandHistory() *CommandHistory {
	return &CommandHistory{
		history: make([]commands.Command, 0),
	}
}

func (h *CommandHistory) Push(cmd commands.Command) {
	h.history = append(h.history, cmd)
}

func (h *CommandHistory) Pop() (commands.Command, error) {
	if len(h.history) == 0 {
		return nil, errors.New("No history found")
	}
	last := h.history[len(h.history)-1]
	h.history = h.history[:len(h.history)-1]
	return last, nil
}
