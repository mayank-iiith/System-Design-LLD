package pricingstrategy

import (
	"math"
	"time"
)

const (
	minuteWisePricing = 1
)

type MinuteWisePricingStrategy struct {
	minuteWisePricing int
}

func NewMinuteWiseParkingStrategy() *MinuteWisePricingStrategy {
	return &MinuteWisePricingStrategy{
		minuteWisePricing: minuteWisePricing,
	}
}

var _ PricingStrategy = (*MinuteWisePricingStrategy)(nil)

// CalculatePrice implements PricingStrategy.
func (ps *MinuteWisePricingStrategy) CalculatePrice(entryTime time.Time) int {
	timeInMinutes := int(math.Ceil(time.Since(entryTime).Minutes()))
	return timeInMinutes * ps.minuteWisePricing
}
