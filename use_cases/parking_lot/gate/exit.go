package gate

import (
	"fmt"
	parkingspotmanager "lld/use_cases/parking_lot/parking_spot_manager"
	"lld/use_cases/parking_lot/ticket"
)

type Exit struct {
	parkingSpotManagerFactory *parkingspotmanager.ParkingSpotManagerFactory
}

func NewExit(f *parkingspotmanager.ParkingSpotManagerFactory) *Exit {
	return &Exit{
		parkingSpotManagerFactory: f,
	}
}

func (e *Exit) ExitVehicle(t *ticket.Ticket) error {
	mgr, err := e.parkingSpotManagerFactory.GetParkingSpotManager(t.Vehicle.GetType())
	if err != nil {
		return err
	}

	cost := t.ParkingSpot.GetPricingStrategy().CalculatePrice(t.EntryTime)
	fmt.Println("Cost for parking is:", cost)

	if err := mgr.EmptyParkingSpot(t.ParkingSpot.SpotNumber()); err != nil {
		return err
	}
	return nil
}
