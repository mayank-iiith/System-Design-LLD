package shape

type Square struct {
	side int
}

func NewSquare(side int) *Square {
	return &Square{
		side: side,
	}
}

func (s *Square) GetType() string {
	return "Square"
}

func (s *Square) Accept(v ShapeVisitor) {
	v.visitForSquare(s)
}
