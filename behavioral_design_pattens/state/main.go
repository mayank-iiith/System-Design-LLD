package main

// State is a behavioral design pattern that lets an object alter its behavior when its internal state changes. It appears as if the object changed its class.

// The State pattern is closely related to the concept of a Finite-State Machine .

// The pattern extracts state-related behaviors into separate state classes and forces the original object to delegate the work to an instance of these classes, instead of acting on its own.

import (
	"fmt"
	vendingmachine "lld/behavioral_design_pattens/state/vending_machine"
	"log"
)

func main() {
	vendingMachine := vendingmachine.NewVendingMachine(1, 10)

	err := vendingMachine.RequestItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println()
	err = vendingMachine.AddItem(2)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.AddItem(3)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println()

	err = vendingMachine.RequestItem()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.InsertMoney(10)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = vendingMachine.DispenseItem()
	if err != nil {
		log.Fatal(err.Error())
	}
}
