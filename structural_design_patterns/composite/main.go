package main

//Composite is a structural design pattern that lets you compose objects into tree structures and then work with these structures as if they were individual objects.

//Composite became a pretty popular solution for the most problems that require building a tree structure. Composite’s great feature is the ability to run methods recursively over the whole tree structure and sum up the results.

import (
	filesystem "lld/structural_design_patterns/composite/file_system"
)

func main() {
	file1 := filesystem.NewFile("File1")
	file2 := filesystem.NewFile("File2")
	file3 := filesystem.NewFile("File3")

	folder1 := filesystem.NewFolder("Folder1")
	folder1.Add(file1)

	folder2 := filesystem.NewFolder("Folder2")
	folder2.Add(file2)
	folder2.Add(file3)
	folder2.Add(folder1)

	folder2.Search("hello")
}
