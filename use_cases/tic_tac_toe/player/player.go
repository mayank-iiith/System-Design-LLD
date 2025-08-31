package player

import "lld/use_cases/tic_tac_toe/piece"

type Player struct {
	name      string
	pieceType piece.PieceType
}

func NewPlayer(name string, pieceType piece.PieceType) *Player {
	return &Player{
		name:      name,
		pieceType: pieceType,
	}
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) PieceType() piece.PieceType {
	return p.pieceType
}
