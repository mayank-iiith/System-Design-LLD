package parkingspot

import (
	pricingstrategy "lld/use_cases/parking_lot/pricing_strategy"
)

type TwoWheelerParkingSpot struct {
	*baseParkingSpot
}

func NewTwoWheelerParkingSpot(spotNumber int, ps pricingstrategy.PricingStrategy) *TwoWheelerParkingSpot {
	return &TwoWheelerParkingSpot{
		&baseParkingSpot{
			spotNumber:      spotNumber,
			pricingStrategy: ps,
		},
	}
}

var _ ParkingSpot = (*TwoWheelerParkingSpot)(nil)
