package main

// Builder is a creational design pattern that lets you construct complex objects step by step. The pattern allows you to produce different types and representations of an object using the same construction code.

import (
	"fmt"
	"lld/creational_design_patterns/builder/director"
	housebuilder "lld/creational_design_patterns/builder/house_builder"
)

func main() {
	normalBuilder := housebuilder.GetBuilder(housebuilder.NormalBuilderType)
	iglooBuilder := housebuilder.GetBuilder(housebuilder.IglooBuilderType)

	dir := director.NewDirector(normalBuilder)
	normalHouse := dir.BuildHouse()
	fmt.Println("Printing Normal House Details:")
	normalHouse.PrintDetails()

	dir.SetBuilder(iglooBuilder)
	iglooHouse := dir.BuildHouse()
	fmt.Println("Printing Igloo House Details:")
	iglooHouse.PrintDetails()
}
