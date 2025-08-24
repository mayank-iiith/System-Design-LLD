package game

import "lld/structural_design_patterns/flyweight/game/dress"

type Game struct {
	terrorists        []*Player
	counterTerrorists []*Player
}

func NewGame() *Game {
	return &Game{
		terrorists:        make([]*Player, 0),
		counterTerrorists: make([]*Player, 0),
	}
}

func (g *Game) AddTerrorist(dressType dress.DressType) error {
	player, err := NewPlayer(TerroristPlayerType, dressType)
	if err != nil {
		return err
	}
	g.terrorists = append(g.terrorists, player)
	return nil
}

func (g *Game) AddCounterTerrorist(dressType dress.DressType) error {
	player, err := NewPlayer(CounterTerroristPlayerType, dressType)
	if err != nil {
		return err
	}
	g.terrorists = append(g.terrorists, player)
	return nil
}
