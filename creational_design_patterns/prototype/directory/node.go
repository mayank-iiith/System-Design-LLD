package directory

type Node interface {
	Print(indent string)
	Clone() Node
}
