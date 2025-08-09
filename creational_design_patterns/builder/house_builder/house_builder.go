package housebuilder

import (
	"lld/creational_design_patterns/builder/house"
)

type Builder interface {
	SetWindowType()
	SetDoorType()
	SetNumFloor()
	GetHouse() *house.House
}

type BuilderType int

const (
	NormalBuilderType BuilderType = iota + 1
	IglooBuilderType
)

func GetBuilder(builderType BuilderType) Builder {
	if builderType == IglooBuilderType {
		return NewIglooBuilder()
	}
	return NewNormalBuilder()
}
