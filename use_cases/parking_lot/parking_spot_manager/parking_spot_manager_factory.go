package parkingspotmanager

import (
	"errors"
	parkingspot "lld/use_cases/parking_lot/parking_spot"
	"lld/use_cases/parking_lot/vehicle"
)

type ParkingSpotManagerFactory struct {
	managers map[parkingspot.ParkingSpotType]ParkingSpotManager
}

func NewParkingSpotManagerFactory() *ParkingSpotManagerFactory {
	return &ParkingSpotManagerFactory{
		managers: make(map[parkingspot.ParkingSpotType]ParkingSpotManager),
	}
}

func (f *ParkingSpotManagerFactory) GetParkingSpotManager(vt vehicle.VehicleType) (ParkingSpotManager, error) {
	pt, err := getCompatibleParkingSpotType(vt)
	if err != nil {
		return nil, err
	}

	mgr, ok := f.managers[pt]
	if !ok {
		mgr, _ = newParkingSpotManager(pt)
		f.managers[pt] = mgr
	}
	return mgr, nil
}

func (f *ParkingSpotManagerFactory) AddParkingSpot(pt parkingspot.ParkingSpotType, spotNumber int) {
	mgr, ok := f.managers[pt]
	if !ok {
		mgr, _ = newParkingSpotManager(pt)
		f.managers[pt] = mgr
	}

	mgr.AddParkingSpot(spotNumber)
}

func getCompatibleParkingSpotType(vt vehicle.VehicleType) (parkingspot.ParkingSpotType, error) {
	switch vt {
	case vehicle.VehicleTypeBike:
		return parkingspot.TwoWheelerParkingSpotType, nil
	case vehicle.VehicleTypeCar:
		return parkingspot.LightFourWheelerParkingSpotType, nil
	case vehicle.VehicleTypeTruck:
		return parkingspot.HeavyFourWheelerParkingSpotType, nil
	default:
		return parkingspot.UnknownParkingSpotType, errors.New("no parking spot type available for vehicle type")
	}
}

func newParkingSpotManager(pt parkingspot.ParkingSpotType) (ParkingSpotManager, error) {
	switch pt {
	case parkingspot.TwoWheelerParkingSpotType:
		return NewTwoWheelerParkingSpotManager(), nil
	case parkingspot.LightFourWheelerParkingSpotType:
		return NewLightFourWheelerParkingSpotManager(), nil
	case parkingspot.HeavyFourWheelerParkingSpotType:
		return NewHeavyFourWheelerParkingSpotManager(), nil
	default:
		return nil, errors.New("unsupported parking spot type")
	}
}
