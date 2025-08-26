package parkingspot

import (
	"errors"
	pricingstrategy "lld/use_cases/parking_lot/pricing_strategy"
	"lld/use_cases/parking_lot/vehicle"
)

type ParkingSpotType int

const (
	TwoWheelerParkingSpotType ParkingSpotType = iota + 1
	LightFourWheelerParkingSpotType
	HeavyFourWheelerParkingSpotType
)

type ParkingSpot interface {
	IsEmpty() bool
	ParkVehicle(vehicle.Vehicle) error
	RemoveVehicle() error
	GetPricingStrategy() pricingstrategy.PricingStrategy
}

type baseParkingSpot struct {
	spotNumber      int
	vehicle         vehicle.Vehicle
	pricingStrategy pricingstrategy.PricingStrategy
}

func (ps *baseParkingSpot) IsEmpty() bool {
	return ps.vehicle == nil
}

func (ps *baseParkingSpot) ParkVehicle(v vehicle.Vehicle) error {
	if !ps.IsEmpty() {
		return errors.New("parking spot is already occupied")
	}
	ps.vehicle = v
	return nil
}

func (ps *baseParkingSpot) RemoveVehicle() error {
	if ps.IsEmpty() {
		return errors.New("parking spot is already empty")
	}
	ps.vehicle = nil
	return nil
}

func (ps *baseParkingSpot) GetPricingStrategy() pricingstrategy.PricingStrategy {
	return ps.pricingStrategy
}
