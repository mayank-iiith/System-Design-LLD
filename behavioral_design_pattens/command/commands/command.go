package commands

import (
	"lld/behavioral_design_pattens/command/editor"
)

// Command defines the contract for all commands
type Command interface {
	Execute() bool
	Undo()
}

// BaseCommand holds common fields and methods for commands
type BaseCommand struct {
	Editor *editor.Editor
	backup string
}

func (c *BaseCommand) Backup() {
	c.backup = c.Editor.Text
}

func (c *BaseCommand) Undo() {
	c.Editor.Text = c.backup
}
