package ticket

import (
	"fmt"
	parkingspot "lld/use_cases/parking_lot/parking_spot"
	"lld/use_cases/parking_lot/vehicle"
	"time"
)

type Ticket struct {
	EntryTime   time.Time
	Vehicle     vehicle.Vehicle
	ParkingSpot parkingspot.ParkingSpot
}

func GenerateTicket(v vehicle.Vehicle, ps parkingspot.ParkingSpot) *Ticket {
	return &Ticket{
		EntryTime:   time.Now(),
		Vehicle:     v,
		ParkingSpot: ps,
	}
}

func (t *Ticket) Print(prefix string) {
	fmt.Println("-------------------------------")
	fmt.Printf("## Parking Ticket for %s ##\n", prefix)
	fmt.Println("EntryTime:", t.EntryTime.Format("2006-01-02 15:04:05"))
	fmt.Println("Vehicle Number:", t.Vehicle.GetVehicleNumber())
	fmt.Println("Parking Spot Number:", t.ParkingSpot.SpotNumber())
	fmt.Println("-------------------------------")
	fmt.Println()
}
