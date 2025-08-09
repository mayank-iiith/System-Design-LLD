package housebuilder

import (
	"lld/creational_design_patterns/builder/house"
)

type IglooBuilder struct {
	windowType house.WindowType
	doorType   house.DoorType
	numFloor   int
}

func NewIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) SetWindowType() {
	b.windowType = house.SnowWindowType
}

func (b *IglooBuilder) SetDoorType() {
	b.doorType = house.SnowDoorType
}

func (b *IglooBuilder) SetNumFloor() {
	b.numFloor = 1
}

func (b *IglooBuilder) GetHouse() *house.House {
	return &house.House{
		WindowType: b.windowType,
		DoorType:   b.doorType,
		NumFloor:   b.numFloor,
	}
}
