package housebuilder

import (
	"lld/creational_design_patterns/builder/house"
)

type NormalBuilder struct {
	windowType house.WindowType
	doorType   house.DoorType
	numFloor   int
}

func NewNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) SetWindowType() {
	b.windowType = house.WoodenWindowType
}

func (b *NormalBuilder) SetDoorType() {
	b.doorType = house.WoodenDoorType
}

func (b *NormalBuilder) SetNumFloor() {
	b.numFloor = 2
}

func (b *NormalBuilder) GetHouse() *house.House {
	return &house.House{
		WindowType: b.windowType,
		DoorType:   b.doorType,
		NumFloor:   b.numFloor,
	}
}
