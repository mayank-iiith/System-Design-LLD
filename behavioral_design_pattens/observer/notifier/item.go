package notifier

import (
	"fmt"
	"lld/behavioral_design_pattens/observer/observer"
)

type Item struct {
	subsribers []observer.Observer
	name       string
	inStock    bool
}

func NewItem(name string) *Item {
	return &Item{
		subsribers: make([]observer.Observer, 0),
		name:       name,
	}
}

func (i *Item) UpdateAvailability() {
	fmt.Printf("Item %s is now in stock\n", i.name)
	i.inStock = true
	i.NotifyAll()
}

func (i *Item) Subscribe(o observer.Observer) {
	i.subsribers = append(i.subsribers, o)
}

func (i *Item) Unsubscribe(o observer.Observer) {
	i.subsribers = removeFromSlice(i.subsribers, o)
}

func (i *Item) NotifyAll() {
	for _, subscriber := range i.subsribers {
		subscriber.Update(i.name)
	}
}

func removeFromSlice(subscribers []observer.Observer, o observer.Observer) []observer.Observer {
	updatedSubscribers := make([]observer.Observer, 0)
	for _, subscriber := range subscribers {
		if subscriber != o {
			updatedSubscribers = append(updatedSubscribers, subscriber)
		}
	}
	return updatedSubscribers
}
