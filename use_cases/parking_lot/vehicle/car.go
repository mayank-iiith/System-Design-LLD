package vehicle

type Car struct {
	vehicleNumber string
}

func NewCar(vehicleNumber string) *Car {
	return &Car{
		vehicleNumber: vehicleNumber,
	}
}

func (c *Car) GetVehicleNumber() string {
	return c.vehicleNumber
}

func (*Car) GetType() VehicleType {
	return VehicleTypeCar
}
