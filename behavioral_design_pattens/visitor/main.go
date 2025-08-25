package main

// Visitor is a behavioral design pattern that lets you separate algorithms from the objects on which they operate.
// It allows adding new behaviors to existing class hierarchy without altering any existing code.

import (
	"fmt"
	"lld/behavioral_design_pattens/visitor/shape"
)

func main() {
	square := shape.NewSquare(2)
	circle := shape.NewCircle(3)
	rectangle := shape.NewRectangle(2, 3)

	areaCalculator := &shape.AreaCalculator{}
	square.Accept(areaCalculator)
	circle.Accept(areaCalculator)
	rectangle.Accept(areaCalculator)

	fmt.Println()

	circumferenceCalculator := &shape.CircumferenceCalculator{}
	square.Accept(circumferenceCalculator)
	circle.Accept(circumferenceCalculator)
	rectangle.Accept(circumferenceCalculator)
}
