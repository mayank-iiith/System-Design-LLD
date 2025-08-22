package railway

import "fmt"

type FreightTrain struct {
	mediator Mediator
}

func NewFreightTrain(mediator Mediator) *FreightTrain {
	return &FreightTrain{
		mediator: mediator,
	}
}

func (t *FreightTrain) Arrive() {
	if !t.mediator.CanArrive(t) {
		fmt.Println("FreightTrain: Arrival Blocked, Waiting...")
		return
	}
	fmt.Println("FreightTrain: Arrived.")
}

func (t *FreightTrain) Depart() {
	fmt.Println("FreightTrain: Leaving...")
	t.mediator.NotifyAboutDeparture()
}

func (t *FreightTrain) PermitArrival() {
	fmt.Println("FreightTrain: Arrival permitted, ariving")
	t.Arrive()
}
