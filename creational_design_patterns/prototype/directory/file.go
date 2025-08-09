package directory

import "fmt"

type File struct {
	name string
}

func NewFile(name string) *File {
	return &File{
		name: name,
	}
}

func (f *File) Print(indent string) {
	fmt.Println(indent + f.name)
}

func (f *File) Clone() Node {
	return NewFile(f.name + "_clone")
}
