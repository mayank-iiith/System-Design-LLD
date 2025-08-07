package sports

import (
	"fmt"
	"lld/creational_design_patterns/abstract_factory_method/adidas"
	"lld/creational_design_patterns/abstract_factory_method/nike"
)

type BrandType int

const (
	AdidasBrandType BrandType = iota + 1
	NikeBrandType
)

func GetSportsFactory(t BrandType) (SportsFactory, error) {
	switch t {
	case AdidasBrandType:
		return &adidas.Adidas{}, nil
	case NikeBrandType:
		return &nike.Nike{}, nil
	default:
		return nil, fmt.Errorf("unsupported brand type")
	}
}
