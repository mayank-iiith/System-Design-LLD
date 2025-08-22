package railway

import "fmt"

type PassengerTrain struct {
	mediator Mediator
}

func NewPassengerTrain(mediator Mediator) *PassengerTrain {
	return &PassengerTrain{
		mediator: mediator,
	}
}

func (t *PassengerTrain) Arrive() {
	if !t.mediator.CanArrive(t) {
		fmt.Println("PassengerTrain: Arrival Blocked, Waiting...")
		return
	}
	fmt.Println("PassengerTrain: Arrived.")
}

func (t *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain: Leaving...")
	t.mediator.NotifyAboutDeparture()
}

func (t *PassengerTrain) PermitArrival() {
	fmt.Println("PassengerTrain: Arrival permitted, ariving")
	t.Arrive()
}
