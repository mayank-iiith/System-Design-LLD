package game

import "lld/structural_design_patterns/flyweight/game/dress"

type PlayerType int

const (
	TerroristPlayerType PlayerType = iota + 1
	CounterTerroristPlayerType
)

type Player struct {
	// Intrinsic State: Dress in the intrinsic state as it can be shared across multiple Terrorist and Counter-Terrorist objects.
	dress      dress.Dress
	playerType PlayerType

	// Extrinsic State: Player location and the player weapon are an extrinsic state as they are different for every object.
	lat  int // latitude
	long int // longitude
}

func NewPlayer(playerType PlayerType, dressType dress.DressType) (*Player, error) {
	dress, err := dress.GetDressFactorySingleInstance().GetDressByType(dressType)
	if err != nil {
		return nil, err
	}

	return &Player{
		playerType: playerType,
		dress:      dress,
	}, nil
}

func (p *Player) NewLocation(lat, long int) {
	p.lat = lat
	p.long = long
}
