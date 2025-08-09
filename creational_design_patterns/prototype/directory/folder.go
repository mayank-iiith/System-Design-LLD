package directory

import "fmt"

type Folder struct {
	name      string
	childrens []Node
}

func NewFolder(name string) *Folder {
	return &Folder{
		name: name,
	}
}

func (f *Folder) AddChildren(node Node) {
	f.childrens = append(f.childrens, node)
}

func (f *Folder) Print(indent string) {
	fmt.Println(indent + f.name)
	for _, child := range f.childrens {
		child.Print(indent + indent)
	}
}

func (f *Folder) Clone() Node {
	clonedFolder := NewFolder(f.name + "_clone")
	for _, child := range f.childrens {
		clonedChild := child.Clone()
		clonedFolder.childrens = append(clonedFolder.childrens, clonedChild)
	}

	return clonedFolder
}
