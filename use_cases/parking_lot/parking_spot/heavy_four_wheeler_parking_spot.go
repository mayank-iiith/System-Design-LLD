package parkingspot

import pricingstrategy "lld/use_cases/parking_lot/pricing_strategy"

type HeavyFourWheelerParkingSpot struct {
	*baseParkingSpot
}

func NewHeavyFourWheelerParkingSpot(spotNumber int, ps pricingstrategy.PricingStrategy) *HeavyFourWheelerParkingSpot {
	return &HeavyFourWheelerParkingSpot{
		&baseParkingSpot{
			spotNumber:      spotNumber,
			pricingStrategy: ps,
		},
	}
}

var _ ParkingSpot = (*HeavyFourWheelerParkingSpot)(nil)
