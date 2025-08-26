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
	ParkVehicle(vehicle.Vehicle) (parkingspot.ParkingSpot, error)
	EmptyParkingSpot(int) error
	FindParkingSpot() (parkingspot.ParkingSpot, error)
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
func (m *baseParkingSpotManager) ParkVehicle(v vehicle.Vehicle) (parkingspot.ParkingSpot, error) {
	ps, err := m.FindParkingSpot()
	if err != nil {
		return nil, err
	}
	err = ps.ParkVehicle(v)
	if err != nil {
		return nil, err
	}
	return ps, nil
}

// EmptyParkingSpot implements ParkingSpotManager.
func (m *baseParkingSpotManager) EmptyParkingSpot(spotNumber int) error {
	for _, ps := range m.parkingSpots {
		if ps.SpotNumber() == spotNumber {
			return ps.RemoveVehicle()
		}
	}
	return fmt.Errorf("parking spot doesn't exists")
}

// FindParkingSpot implements ParkingSpotManager.
func (m *baseParkingSpotManager) FindParkingSpot() (parkingspot.ParkingSpot, error) {
	for _, ps := range m.parkingSpots {
		if ps.IsEmpty() {
			return ps, nil
		}
	}
	return nil, errors.New("no parking spots available")
}
