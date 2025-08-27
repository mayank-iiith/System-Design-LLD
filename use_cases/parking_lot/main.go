package main

import (
	"fmt"
	"lld/use_cases/parking_lot/gate"
	parkingspot "lld/use_cases/parking_lot/parking_spot"
	parkingspotmanager "lld/use_cases/parking_lot/parking_spot_manager"
	"lld/use_cases/parking_lot/vehicle"
)

func main() {
	bike1 := vehicle.NewBike("b123")
	bike2 := vehicle.NewBike("b456")

	car1 := vehicle.NewCar("c123")
	car2 := vehicle.NewCar("c456")

	truck1 := vehicle.NewTruck("t456")

	parkingSpotManagerFactory := parkingspotmanager.NewParkingSpotManagerFactory()
	addParkingSpots(parkingSpotManagerFactory, parkingspot.TwoWheelerParkingSpotType, 4)
	addParkingSpots(parkingSpotManagerFactory, parkingspot.LightFourWheelerParkingSpotType, 1)
	addParkingSpots(parkingSpotManagerFactory, parkingspot.HeavyFourWheelerParkingSpotType, 3)

	entranceGate := gate.NewEntrance(parkingSpotManagerFactory)
	exitGate := gate.NewExit(parkingSpotManagerFactory)

	ticketForBike1, err := entranceGate.BookSpot(bike1)
	if err != nil {
		fmt.Println(err)
	}
	ticketForBike1.Print("Bike1")

	ticketForBike2, err := entranceGate.BookSpot(bike2)
	if err != nil {
		fmt.Println(err)
	}
	ticketForBike2.Print("Bike2")

	ticketForTruck1, err := entranceGate.BookSpot(truck1)
	if err != nil {
		fmt.Println(err)
	}
	ticketForTruck1.Print("Truck1")

	ticketForCar1, err := entranceGate.BookSpot(car1)
	if err != nil {
		fmt.Println(err)
	}
	ticketForCar1.Print("Car1")

	_ = exitGate.ExitVehicle(ticketForBike1)
	_ = exitGate.ExitVehicle(ticketForBike2)
	_ = exitGate.ExitVehicle(ticketForTruck1)

	_, err = entranceGate.BookSpot(car2)
	if err != nil {
		fmt.Println(err)
	}

	_ = exitGate.ExitVehicle(ticketForCar1)

	ticketForCar2, err := entranceGate.BookSpot(car2)
	if err != nil {
		fmt.Println(err)
	}
	ticketForCar2.Print("Car2")

	_ = exitGate.ExitVehicle(ticketForCar2)
}

func addParkingSpots(
	f *parkingspotmanager.ParkingSpotManagerFactory,
	pt parkingspot.ParkingSpotType,
	numOfParkingSpots int,
) {
	for i := 1; i <= numOfParkingSpots; i++ {
		f.AddParkingSpot(pt, i)
	}
}
