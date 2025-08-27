package vehicle

type VehicleType int

const (
	VehicleTypeBike VehicleType = iota + 1
	VehicleTypeCar
	VehicleTypeTruck
)

type Vehicle interface {
	GetVehicleNumber() string
	GetType() VehicleType
}
