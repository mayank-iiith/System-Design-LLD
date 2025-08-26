package parkingspot

import (
	pricingstrategy "lld/use_cases/parking_lot/pricing_strategy"
)

type LightFourWheelerParkingSpot struct {
	*baseParkingSpot
}

func NewLightFourWheelerParkingSpot(spotNumber int, ps pricingstrategy.PricingStrategy) *LightFourWheelerParkingSpot {
	return &LightFourWheelerParkingSpot{
		&baseParkingSpot{
			spotNumber:      spotNumber,
			pricingStrategy: ps,
		},
	}
}

var _ ParkingSpot = (*LightFourWheelerParkingSpot)(nil)
