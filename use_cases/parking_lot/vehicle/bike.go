package vehicle

type Bike struct {
	vehicleNumber string
}

func NewBike(vehicleNumber string) *Bike {
	return &Bike{
		vehicleNumber: vehicleNumber,
	}
}

func (b *Bike) GetVehicleNumber() string {
	return b.vehicleNumber
}

func (*Bike) GetType() VehicleType {
	return VehicleTypeBike
}
