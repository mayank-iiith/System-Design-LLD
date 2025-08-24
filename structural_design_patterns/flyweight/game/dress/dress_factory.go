package dress

import (
	"errors"
	"fmt"
)

type DressType int

const (
	TerroristDressType DressType = iota + 1
	CounterTerroristDressType
)

type dressFactory struct {
	dressMap map[DressType]Dress
}

var (
	dressFactorySingleInstance = &dressFactory{
		dressMap: make(map[DressType]Dress),
	}
)

func GetDressFactorySingleInstance() *dressFactory {
	return dressFactorySingleInstance
}

func (d *dressFactory) GetDressByType(dressType DressType) (Dress, error) {
	if dress, ok := d.dressMap[dressType]; ok {
		fmt.Println("Dress found with dressType:", dressType)
		return dress, nil
	}

	switch dressType {
	case TerroristDressType:
		d.dressMap[dressType] = newTerroristDress()
		fmt.Println("New Dress created with dressType:", dressType)
		return d.dressMap[dressType], nil

	case CounterTerroristDressType:
		fmt.Println("New Dress created with dressType:", dressType)
		d.dressMap[dressType] = newCounterTerroristDress()
		return d.dressMap[dressType], nil

	default:
		return nil, errors.New("dress type not supported")
	}
}
