package main

// Abstract Factory is a creational design pattern, which solves the problem of creating entire product families without specifying their concrete classes.

import (
	"fmt"
	"log"

	"lld/creational_design_patterns/abstract_factory_method/sports"
)

func main() {
	//sportFactory, err := sports.GetSportsFactory(sports.BrandType(5))
	//sportFactory, err := sports.GetSportsFactory(sports.NikeBrandType)

	sportFactory, err := sports.GetSportsFactory(sports.AdidasBrandType)
	if err != nil {
		log.Fatalf("unable to get sport factory, error: %+v", err)
	}
	shoe := sportFactory.MakeShoe(18)
	fmt.Println(shoe.GetLogoAndType())

	shirt := sportFactory.MakeShirt(14)
	fmt.Println(shirt.GetLogoAndType())
}
