package pricingstrategy

import (
	"math"
	"time"
)

const (
	hourlyPriceForFirstHour      = 50
	hourlyPriceForSubsequentHour = 50
)

type DynamicHourlyPricingStrategy struct {
	hourlyPriceForFirstHour      int
	hourlyPriceForSubsequentHour int
}

func NewDynamicHourlyPricingStrategy() *DynamicHourlyPricingStrategy {
	return &DynamicHourlyPricingStrategy{
		hourlyPriceForFirstHour:      hourlyPriceForFirstHour,
		hourlyPriceForSubsequentHour: hourlyPriceForSubsequentHour,
	}
}

var _ PricingStrategy = (*DynamicHourlyPricingStrategy)(nil)

// CalculatePrice implements PricingStrategy.
func (ps *DynamicHourlyPricingStrategy) CalculatePrice(entryTime time.Time) int {
	timeInHours := int(math.Ceil(time.Since(entryTime).Hours()))
	return ps.hourlyPriceForFirstHour + ps.hourlyPriceForSubsequentHour*(timeInHours-1)
}
