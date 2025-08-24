package commands

type CopyCommand struct {
	BaseCommand
}

func (c *CopyCommand) Execute() bool {
	c.Editor.Clipboard = c.Editor.GetSelection()
	return false // No change to editor content
}
