package parkingspotmanager

import (
	parkingspot "lld/use_cases/parking_lot/parking_spot"
	pricingstrategy "lld/use_cases/parking_lot/pricing_strategy"
)

type LightFourWheelerParkingSpotManager struct {
	*baseParkingSpotManager
}

func NewLightFourWheelerParkingSpotManager(spotNumber int, ps pricingstrategy.PricingStrategy) *LightFourWheelerParkingSpotManager {
	return &LightFourWheelerParkingSpotManager{
		&baseParkingSpotManager{
			parkingSpots:    make([]parkingspot.ParkingSpot, 0),
			pricingStrategy: ps,
		},
	}
}

var _ ParkingSpotManager = (*LightFourWheelerParkingSpotManager)(nil)

// AddParkingSpot implements ParkingSpotManager.
func (m *LightFourWheelerParkingSpotManager) AddParkingSpot(spotNumber int) {
	m.parkingSpots = append(m.parkingSpots, parkingspot.NewLightFourWheelerParkingSpot(spotNumber, m.pricingStrategy))
}
