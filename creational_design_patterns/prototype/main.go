package main

import (
	"fmt"
	"lld/creational_design_patterns/prototype/directory"
)

const (
	indent = "  "
)

func main() {
	file1 := directory.NewFile("file1")
	file2 := directory.NewFile("file2")
	file3 := directory.NewFile("file3")

	folder1 := directory.NewFolder("folder1")
	folder1.AddChildren(file1)

	folder2 := directory.NewFolder("folder2")
	folder2.AddChildren(folder1)
	folder2.AddChildren(file2)
	folder2.AddChildren(file3)

	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print(indent)

	folder2Clone := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone of Folder2")
	folder2Clone.Print(indent)
}
