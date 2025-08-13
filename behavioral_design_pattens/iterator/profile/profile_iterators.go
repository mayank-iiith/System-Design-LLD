package profile

type ProfileIterator interface {
	HasNext() bool
	GetNext() (*Profile, error)
	Reset()
}
