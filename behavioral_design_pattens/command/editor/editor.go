package editor

type Editor struct {
	Text      string
	Clipboard string
}

func NewEditor() *Editor {
	return &Editor{}
}

func (e *Editor) GetSelection() string {
	return e.Text
}

func (e *Editor) DeleteSelection() {
	e.Text = ""
}

func (e *Editor) ReplaceSelection(text string) {
	e.Text = text
}
