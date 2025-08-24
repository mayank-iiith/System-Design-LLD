package commands

type CutCommand struct {
	BaseCommand
}

func (c *CutCommand) Execute() bool {
	c.Backup()
	c.Editor.Clipboard = c.Editor.GetSelection()
	c.Editor.DeleteSelection()
	return true // Editor content changed
}
