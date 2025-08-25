package shape

import "fmt"

type AreaCalculator struct {
}

// AreaCalculator Implements ShapeVisitor
var _ ShapeVisitor = (*AreaCalculator)(nil)

// visitForCircle implements ShapeVisitor.
func (a *AreaCalculator) visitForCircle(c *Circle) {
	area := 3.14 * float64(c.radius) * float64(c.radius)
	fmt.Println("Area of Circle:", area)
}

// visitForRectangle implements ShapeVisitor.
func (a *AreaCalculator) visitForRectangle(r *Rectangle) {
	area := r.length * r.breadth
	fmt.Println("Area of Rectangle:", area)
}

// visitForSquare implements ShapeVisitor.
func (a *AreaCalculator) visitForSquare(s *Square) {
	area := s.side * s.side
	fmt.Println("Area of Square:", area)
}
