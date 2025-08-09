package house

import "fmt"

type WindowType string

const (
	WoodenWindowType WindowType = "Wooden Window"
	SnowWindowType   WindowType = "Snow Window"
)

func (t WindowType) string() string {
	return string(t)
}

type DoorType string

const (
	WoodenDoorType DoorType = "Wooden Door"
	SnowDoorType   DoorType = "Snow Door"
)

func (t DoorType) string() string {
	return string(t)
}

type House struct {
	WindowType WindowType
	DoorType   DoorType
	NumFloor   int
}

func (h *House) PrintDetails() {
	fmt.Println("Window Type:", h.WindowType.string())
	fmt.Println("DoorType Type:", h.DoorType.string())
	fmt.Println("Number of Floor:", h.NumFloor)
}
