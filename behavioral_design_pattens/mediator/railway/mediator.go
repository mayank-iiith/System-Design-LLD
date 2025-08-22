package railway

type Mediator interface {
	CanArrive(Train) bool
	NotifyAboutDeparture()
}
