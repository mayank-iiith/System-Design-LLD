package shape

type Circle struct {
	radius int
}

func NewCircle(radius int) *Circle {
	return &Circle{
		radius: radius,
	}
}

func (c *Circle) GetType() string {
	return "Circle"
}

func (c *Circle) Accept(v ShapeVisitor) {
	v.visitForCircle(c)
}
