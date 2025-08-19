package filesystem

import "fmt"

type File struct {
	name string
}

func NewFile(name string) *File {
	return &File{
		name: name,
	}
}

func (f *File) Search(keyword string) {
	fmt.Printf("Searching for keyword %q in file %q\n", keyword, f.name)
}
