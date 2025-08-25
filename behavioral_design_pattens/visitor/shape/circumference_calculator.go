package shape

import (
	"fmt"
	"math"
)

type CircumferenceCalculator struct {
}

// CircumferenceCalculator Implements ShapeVisitor
var _ ShapeVisitor = (*CircumferenceCalculator)(nil)

// visitForCircle implements ShapeVisitor.
func (c *CircumferenceCalculator) visitForCircle(ci *Circle) {
	circumference := 2 * math.Pi * float64(ci.radius)
	fmt.Println("Circumference of Circle:", circumference)
}

// visitForRectangle implements ShapeVisitor.
func (c *CircumferenceCalculator) visitForRectangle(r *Rectangle) {
	circumference := 2 * (r.length + r.breadth)
	fmt.Println("Circumference of Rectangle:", circumference)
}

// visitForSquare implements ShapeVisitor.
func (c *CircumferenceCalculator) visitForSquare(s *Square) {
	circumference := 4 * s.side
	fmt.Println("Circumference of Square:", circumference)
}
