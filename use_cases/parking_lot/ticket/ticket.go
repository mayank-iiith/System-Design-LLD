package ticket

import (
	parkingspot "lld/use_cases/parking_lot/parking_spot"
	"lld/use_cases/parking_lot/ticket"
	"lld/use_cases/parking_lot/vehicle"
	"time"
)

type Ticket struct {
	EntryTime   time.Time
	Vehicle     vehicle.Vehicle
	ParkingSpot parkingspot.ParkingSpot
}

func GenerateTicket(v vehicle.Vehicle, ps parkingspot.ParkingSpot) *Ticket {
	return &ticket.Ticket{
		EntryTime:   time.Now(),
		Vehicle:     v,
		ParkingSpot: ps,
	}
}
