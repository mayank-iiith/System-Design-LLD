package director

import (
	"lld/creational_design_patterns/builder/house"
	housebuilder "lld/creational_design_patterns/builder/house_builder"
)

type Director struct {
	Builder housebuilder.Builder
}

func NewDirector(b housebuilder.Builder) *Director {
	return &Director{
		Builder: b,
	}
}

func (d *Director) SetBuilder(b housebuilder.Builder) {
	d.Builder = b
}

func (d *Director) BuildHouse() *house.House {
	d.Builder.SetWindowType()
	d.Builder.SetDoorType()
	d.Builder.SetNumFloor()
	return d.Builder.GetHouse()
}
