package main

// Flyweight is a structural design pattern that lets you fit more objects into the available amount of RAM by sharing common parts of state between multiple objects instead of keeping all of the data in each object.
// This pattern keeps the memory consumption low.
// Ref: https://refactoring.guru/design-patterns/flyweight/go/example#example-0

import (
	"lld/structural_design_patterns/flyweight/game"
	"lld/structural_design_patterns/flyweight/game/dress"
)

func main() {
	game1 := game.NewGame()

	// Add Terrorist
	_ = game1.AddTerrorist(dress.TerroristDressType)
	_ = game1.AddTerrorist(dress.TerroristDressType)
	_ = game1.AddTerrorist(dress.TerroristDressType)
	_ = game1.AddTerrorist(dress.TerroristDressType)

	// Add CounterTerrorist
	_ = game1.AddCounterTerrorist(dress.CounterTerroristDressType)
	_ = game1.AddCounterTerrorist(dress.CounterTerroristDressType)
	_ = game1.AddCounterTerrorist(dress.CounterTerroristDressType)
}
