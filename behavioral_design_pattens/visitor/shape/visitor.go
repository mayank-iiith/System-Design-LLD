package shape

type ShapeVisitor interface {
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
	visitForSquare(*Square)
}
