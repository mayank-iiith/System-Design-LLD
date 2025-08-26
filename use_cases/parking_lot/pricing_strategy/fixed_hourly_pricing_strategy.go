package pricingstrategy

import (
	"math"
	"time"
)

const (
	hourlyPrice = 40
)

type FixedHourlyPricingStrategy struct {
	hourlyPrice int
}

func NewFixedHourlyPricingStrategy() *FixedHourlyPricingStrategy {
	return &FixedHourlyPricingStrategy{
		hourlyPrice: hourlyPrice,
	}
}

var _ PricingStrategy = (*FixedHourlyPricingStrategy)(nil)

// CalculatePrice implements PricingStrategy.
func (ps *FixedHourlyPricingStrategy) CalculatePrice(entryTime time.Time) int {
	timeInHours := int(math.Ceil(time.Since(entryTime).Hours()))
	return timeInHours * ps.hourlyPrice
}
