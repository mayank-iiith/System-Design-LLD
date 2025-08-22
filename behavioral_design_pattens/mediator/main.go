package main

import (
	"lld/behavioral_design_pattens/mediator/railway"
)

// Mediator is a behavioral design pattern that lets you reduce chaotic dependencies between objects. The pattern restricts direct communications between the objects and forces them to collaborate only via a mediator object.

// The Mediator makes it easy to modify, extend and reuse individual components because they’re no longer dependent on the dozens of other classes.

func main() {
	stationManager := railway.NewStationManager()

	passengerTrain1 := railway.NewPassengerTrain(stationManager)
	passengerTrain2 := railway.NewPassengerTrain(stationManager)
	freightTrain := railway.NewFreightTrain(stationManager)

	passengerTrain1.Arrive()
	freightTrain.Arrive()
	passengerTrain2.Arrive()

	passengerTrain1.Depart()
	freightTrain.Depart()
	passengerTrain2.Depart()
}
