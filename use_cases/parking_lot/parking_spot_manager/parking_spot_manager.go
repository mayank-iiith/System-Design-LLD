package parkingspotmanager

import (
	"errors"
	"fmt"
	parkingspot "lld/use_cases/parking_lot/parking_spot"
	pricingstrategy "lld/use_cases/parking_lot/pricing_strategy"
	"lld/use_cases/parking_lot/vehicle"
	"slices"
)

type ParkingSpotManager interface {
	AddParkingSpot(int)
	RemoveParkingSpot(int)
	ParkVehicle(vehicle.Vehicle) error
	EmptyParkingSlot(int) error
	findParkingSpot() (parkingspot.ParkingSpot, error)
}

type baseParkingSpotManager struct {
	parkingSpots    []parkingspot.ParkingSpot
	pricingStrategy pricingstrategy.PricingStrategy
}

// RemoveParkingSpot implements ParkingSpotManager.
func (m *baseParkingSpotManager) RemoveParkingSpot(spotNumber int) {
	m.parkingSpots = slices.DeleteFunc(m.parkingSpots, func(p parkingspot.ParkingSpot) bool {
		return p.SpotNumber() == spotNumber
	})
}

// ParkVehicle implements ParkingSpotManager.
func (m *baseParkingSpotManager) ParkVehicle(v vehicle.Vehicle) error {
	ps, err := m.findParkingSpot()
	if err != nil {
		return err
	}
	return ps.ParkVehicle(v)
}

// EmptyParkingSlot implements ParkingSpotManager.
func (m *baseParkingSpotManager) EmptyParkingSlot(spotNumber int) error {
	for _, ps := range m.parkingSpots {
		if ps.SpotNumber() == spotNumber {
			return ps.RemoveVehicle()
		}
	}
	return fmt.Errorf("parking spot doesn't exists")
}

// findParkingSpot implements ParkingSpotManager.
func (m *baseParkingSpotManager) findParkingSpot() (parkingspot.ParkingSpot, error) {
	for _, ps := range m.parkingSpots {
		if ps.IsEmpty() {
			return ps, nil
		}
	}
	return nil, errors.New("no parking spots available")
}
