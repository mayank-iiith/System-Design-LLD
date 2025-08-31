package player

import (
	"lld/use_cases/snake_and_ladder/piece"
)

type Player struct {
	name  string
	piece *piece.Piece
}

func NewPlayer(name string) *Player {
	return &Player{
		name:  name,
		piece: piece.NewPiece(),
	}
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Position() int {
	return p.piece.Position()
}

func (p *Player) SetPosition(pos int) {
	p.piece.SetPosition(pos)

}
