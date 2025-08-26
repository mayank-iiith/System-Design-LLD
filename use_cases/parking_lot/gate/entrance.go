package gate

import (
	parkingspotmanager "lld/use_cases/parking_lot/parking_spot_manager"
	"lld/use_cases/parking_lot/ticket"
	"lld/use_cases/parking_lot/vehicle"
)

type Entrance struct {
	parkingSpotManagerFactory parkingspotmanager.ParkingSpotManagerFactory
}

func NewEntrance(f parkingspotmanager.ParkingSpotManagerFactory) *Entrance {
	return &Entrance{
		parkingSpotManagerFactory: f,
	}
}

func (e *Entrance) BookSpot(v vehicle.Vehicle) (*ticket.Ticket, error) {
	mgr, err := e.parkingSpotManagerFactory.GetParkingSpotManager(v.GetType())
	if err != nil {
		return nil, err
	}

	ps, err := mgr.ParkVehicle(v)
	if err != nil {
		return nil, err
	}

	return ticket.GenerateTicket(v, ps), nil
}
