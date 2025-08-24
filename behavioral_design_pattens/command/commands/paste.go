package commands

type PasteCommand struct {
	BaseCommand
}

func (c *PasteCommand) Execute() bool {
	c.Backup()
	c.Editor.ReplaceSelection(c.Editor.Clipboard)
	return false // Editor content changed
}
