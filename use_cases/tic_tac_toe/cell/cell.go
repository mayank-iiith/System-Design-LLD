package cell

import (
	"lld/use_cases/tic_tac_toe/piece"
)

type Cell struct {
	pieceType piece.PieceType
}

func NewCell() *Cell {
	return &Cell{
		pieceType: piece.PieceTypeEmpty,
	}
}

func (c *Cell) HasPiece() bool {
	return c.pieceType != piece.PieceTypeEmpty
}

func (c *Cell) SetPiece(pt piece.PieceType) {
	c.pieceType = pt
}

func (c *Cell) GetPiece() piece.PieceType {
	return c.pieceType
}
