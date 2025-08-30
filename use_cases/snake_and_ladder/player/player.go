package player

import (
	"fmt"
	"lld/use_cases/snake_and_ladder/board"
	"lld/use_cases/snake_and_ladder/dice"
	"lld/use_cases/snake_and_ladder/piece"
)

type Player struct {
	name  string
	piece *piece.Piece
	dice  *dice.Dice
	board *board.Board
}

func NewPlayer(name string, dice *dice.Dice, board *board.Board) *Player {
	return &Player{
		name:  name,
		piece: piece.NewPiece(),
		dice:  dice,
		board: board,
	}
}

func (p *Player) RollDiceAndMove() {
	move := p.dice.Roll()
	currentPosition := p.piece.GetCurrentPosition()
	newPosition, err := p.board.Move(currentPosition, move)
	if err != nil {
		fmt.Printf("%s rolled a %d but cannot move from %d as it exceeds board limit, total gain: %d\n", p.name, move, currentPosition, 0)
		return
	}
	p.piece.SetNewPosition(newPosition)
	fmt.Printf("%s rolled a %d and moved from %d to %d, total gain: %d\n", p.name, move, currentPosition, newPosition, newPosition-currentPosition)
	if p.board.IsOver() {
		fmt.Printf("%s has won the game!\n", p.name)
	}
}
