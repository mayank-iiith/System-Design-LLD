package main

import (
	"lld/behavioral_design_pattens/observer/notifier"
	"lld/behavioral_design_pattens/observer/observer"
)

// Observer is a behavioral design pattern that lets you define a subscription mechanism to notify multiple objects about any events that happen to the object they’re observing.

func main() {
	shirtItem := notifier.NewItem("Nike Shirt")

	ob1 := observer.NewCustomer("abc@example.com")
	ob2 := observer.NewCustomer("pqr@example.com")
	ob3 := observer.NewCustomer("xyz@example.com")

	shirtItem.Subscribe(ob1)
	shirtItem.Subscribe(ob2)
	shirtItem.Subscribe(ob3)

	shirtItem.UpdateAvailability()
}
