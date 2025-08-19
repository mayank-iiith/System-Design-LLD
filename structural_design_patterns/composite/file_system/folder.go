package filesystem

import "fmt"

type Folder struct {
	name   string
	childs []FileSystem
}

func NewFolder(name string) *Folder {
	return &Folder{
		name:   name,
		childs: make([]FileSystem, 0),
	}
}

func (f *Folder) Add(child FileSystem) {
	f.childs = append(f.childs, child)
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Serching recursively for keyword %q in folder %q\n", keyword, f.name)
	for _, child := range f.childs {
		child.Search(keyword)
	}
}
