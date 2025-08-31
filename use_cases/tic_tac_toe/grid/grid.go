package grid

import (
	"fmt"
	"lld/use_cases/tic_tac_toe/cell"
	"lld/use_cases/tic_tac_toe/piece"
)

type Grid struct {
	dimension int
	cells     [][]*cell.Cell
}

func NewGrid(dimension int) *Grid {
	cells := make([][]*cell.Cell, dimension)
	for i := range dimension {
		cellRow := make([]*cell.Cell, dimension)
		for j := range dimension {
			cellRow[j] = cell.NewCell()
		}
		cells[i] = cellRow
	}

	g := &Grid{
		dimension: dimension,
		cells:     cells,
	}

	g.Print()

	return g
}

func (g *Grid) Print() {
	for i := range g.dimension {
		for j := range g.dimension {
			fmt.Printf("%s ", g.cells[i][j].GetPiece().String())
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *Grid) GetCell(posX, posY int) *cell.Cell {
	return g.cells[posX-1][posY-1]
}

func (g *Grid) IsValidPosition(posX, posY int) bool {
	return posX >= 1 && posX <= g.dimension &&
		posY >= 1 && posY <= g.dimension
}

func (g *Grid) HasPiece(posX, posY int) bool {
	return g.GetCell(posX, posY).HasPiece()
}

// PlacePiece - Place the piece, and return over, win
func (g *Grid) PlacePiece(pt piece.PieceType, posX, posY int) (bool, bool) {
	g.GetCell(posX, posY).SetPiece(pt)
	if g.winnerFound(posX, posY) {
		return true, true
	} else if g.hasAllCellOccupied() {
		return true, false
	}
	return false, false
}

func (g *Grid) winnerFound(posX, posY int) bool {
	// Check row, col, diagonal1 and diagonal2
	rowWin := true
	for j := range g.dimension {
		if g.GetCell(posX, j+1).GetPiece() != g.GetCell(posX, posY).GetPiece() {
			rowWin = false
			break
		}
	}

	if rowWin {
		return true
	}

	colWin := true
	for i := range g.dimension {
		if g.GetCell(i+1, posY).GetPiece() != g.GetCell(posX, posY).GetPiece() {
			colWin = false
			break
		}
	}

	if colWin {
		return true
	}

	diag1Win := false
	if posX == posY {
		diag1Win = true
		for i := range g.dimension {
			if g.GetCell(i+1, i+1).GetPiece() != g.GetCell(posX, posY).GetPiece() {
				diag1Win = false
				break
			}
		}
	}

	if diag1Win {
		return true
	}

	diag2Win := false
	if posX+posY == g.dimension+1 {
		diag2Win = true
		for i := range g.dimension {
			if g.GetCell(i+1, g.dimension-i).GetPiece() != g.GetCell(posX, posY).GetPiece() {
				diag2Win = false
				break
			}
		}
	}

	if diag2Win {
		return true
	}

	return false
}

func (g *Grid) hasAllCellOccupied() bool {
	for i := range g.dimension {
		for j := range g.dimension {
			if !g.cells[i][j].HasPiece() {
				return false
			}
		}
	}
	return true
}
