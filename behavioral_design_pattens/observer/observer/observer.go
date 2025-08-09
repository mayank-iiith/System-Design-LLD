package observer

type Observer interface {
	Update(itemName string)
	GetId() string
}
