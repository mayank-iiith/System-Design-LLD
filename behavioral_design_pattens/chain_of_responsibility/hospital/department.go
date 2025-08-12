package hospital

type Department interface {
	SetNext(Department)
	Execute(*Patient)
}
