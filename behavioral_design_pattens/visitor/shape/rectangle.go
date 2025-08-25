package shape

type Rectangle struct {
	length, breadth int
}

func NewRectangle(length, breadth int) *Rectangle {
	return &Rectangle{
		length:  length,
		breadth: breadth,
	}
}

func (r *Rectangle) GetType() string {
	return "Rectangle"
}

func (r *Rectangle) Accept(v ShapeVisitor) {
	v.visitForRectangle(r)
}
