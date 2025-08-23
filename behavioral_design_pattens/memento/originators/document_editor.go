package originators

import "lld/behavioral_design_pattens/memento/types"

// DocumentEditor is a concrete originator
type DocumentEditor struct {
	content string
}

// DocumentEditorMemento is the snapshot of DocumentEditor
type DocumentEditorMemento struct {
	Content string
}

func NewDocumentEditor() *DocumentEditor {
	return &DocumentEditor{}
}

func (e *DocumentEditor) Type(text string) {
	e.content += text
}

func (e *DocumentEditor) Content() string {
	return e.content
}

// Implement Originator interface
var _ types.Originator = (*DocumentEditor)(nil)

// TakeSnapshot implements types.Originator.
func (e *DocumentEditor) TakeSnapshot() types.Memento {
	return DocumentEditorMemento{
		Content: e.content,
	}
}

// Restore implements types.Originator.
func (e *DocumentEditor) Restore(m types.Memento) {
	if snap, ok := m.(DocumentEditorMemento); ok {
		e.content = snap.Content
	}
}
