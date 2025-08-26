package parkingspotmanager

import (
	parkingspot "lld/use_cases/parking_lot/parking_spot"
	pricingstrategy "lld/use_cases/parking_lot/pricing_strategy"
)

type TwoWheelerParkingSpotManager struct {
	*baseParkingSpotManager
}

func NewTwoWheelerParkingSpotManager(spotNumber int, ps pricingstrategy.PricingStrategy) *TwoWheelerParkingSpotManager {
	return &TwoWheelerParkingSpotManager{
		&baseParkingSpotManager{
			parkingSpots:    make([]parkingspot.ParkingSpot, 0),
			pricingStrategy: ps,
		},
	}
}

var _ ParkingSpotManager = (*TwoWheelerParkingSpotManager)(nil)

// AddParkingSpot implements ParkingSpotManager.
func (m *TwoWheelerParkingSpotManager) AddParkingSpot(spotNumber int) {
	m.parkingSpots = append(m.parkingSpots, parkingspot.NewTwoWheelerParkingSpot(spotNumber, m.pricingStrategy))
}
