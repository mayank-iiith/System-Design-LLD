package main

// Facade is a structural design pattern that provides a simplified (but limited) interface to a complex system of classes, library or framework.
// While Facade decreases the overall complexity of the application, it also helps to move unwanted dependencies to one place.

import (
	"fmt"
	"lld/structural_design_patterns/facade/payment"
	"log"
)

func main() {
	walletFacade := payment.NewWalletFacade("abc", 1234)

	//fmt.Println()
	//if err := walletFacade.AddMonetToWallet("def", 1234, 10); err != nil {
	//	log.Fatalf("Error: %s\n", err.Error())
	//}

	//fmt.Println()
	//if err := walletFacade.AddMonetToWallet("abc", 3456, 10); err != nil {
	//	log.Fatalf("Error: %s\n", err.Error())
	//}

	fmt.Println()
	if err := walletFacade.AddMonetToWallet("abc", 1234, 10); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	if err := walletFacade.DeductMonetFromWallet("abc", 1234, 5); err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	//fmt.Println()
	//if err := walletFacade.DeductMonetFromWallet("abc", 1234, 15); err != nil {
	//	log.Fatalf("Error: %s\n", err.Error())
	//}
}
