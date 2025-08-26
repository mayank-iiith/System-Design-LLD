package vehicle

type Truck struct {
	vehicleNumber string
}

func NewTruck(vehicleNumber string) *Truck {
	return &Truck{
		vehicleNumber: vehicleNumber,
	}
}

func (t *Truck) GetVehicleNumber() string {
	return t.vehicleNumber
}

func (*Truck) GetType() VehicleType {
	return VehicleTypeTruck
}
