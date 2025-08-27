package pricingstrategy

import "time"

type PricingStrategyType int

const (
	MinuteWisePricingStrategyType = iota + 1
	FixedHourlyPricingStrategyType
	DynamicHourlyPricingStrategyType
)

type PricingStrategy interface {
	CalculatePrice(entryTime time.Time) int
}
