package main

// Chain of Responsibility is a behavioral design pattern that lets you pass requests along a chain of handlers. Upon receiving a request, each handler decides either to process the request or to pass it to the next handler in the chain.
// Ref 1: https://refactoring.guru/design-patterns/chain-of-responsibility
// Ref 2: https://refactoring.guru/design-patterns/chain-of-responsibility/go/example#example-0

import (
	"lld/behavioral_design_pattens/chain_of_responsibility/hospital"
)

func main() {
	cashier := &hospital.Cashier{}

	//Set next for medical department
	medical := &hospital.Medical{}
	medical.SetNext(cashier)

	//Set next for doctor department
	doctor := &hospital.Doctor{}
	doctor.SetNext(medical)

	//Set next for reception department
	reception := hospital.Reception{}
	reception.SetNext(doctor)

	patient := &hospital.Patient{Name: "abc"}

	//Patient visiting
	reception.Execute(patient)
}
