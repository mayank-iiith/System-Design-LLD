package parkingspotmanager

import (
	parkingspot "lld/use_cases/parking_lot/parking_spot"
	pricingstrategy "lld/use_cases/parking_lot/pricing_strategy"
)

type HeavyFourWheelerParkingSpotManager struct {
	*baseParkingSpotManager
}

func NewHeavyFourWheelerParkingSpotManager(spotNumber int, ps pricingstrategy.PricingStrategy) *HeavyFourWheelerParkingSpotManager {
	return &HeavyFourWheelerParkingSpotManager{
		&baseParkingSpotManager{
			parkingSpots:    make([]parkingspot.ParkingSpot, 0),
			pricingStrategy: ps,
		},
	}
}

var _ ParkingSpotManager = (*HeavyFourWheelerParkingSpotManager)(nil)

// AddParkingSpot implements ParkingSpotManager.
func (m *HeavyFourWheelerParkingSpotManager) AddParkingSpot(spotNumber int) {
	m.parkingSpots = append(m.parkingSpots, parkingspot.NewHeavyFourWheelerParkingSpot(spotNumber, m.pricingStrategy))
}
