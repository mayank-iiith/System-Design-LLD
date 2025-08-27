package pricingstrategy

import "errors"

func NewParkingStrategy(t PricingStrategyType) (PricingStrategy, error) {
	switch t {
	case MinuteWisePricingStrategyType:
		return NewMinuteWiseParkingStrategy(), nil
	case FixedHourlyPricingStrategyType:
		return NewFixedHourlyPricingStrategy(), nil
	case DynamicHourlyPricingStrategyType:
		return NewDynamicHourlyPricingStrategy(), nil
	default:
		return nil, errors.New("invalid pricing strategy type")
	}
}
