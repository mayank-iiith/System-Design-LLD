package board

import (
	"errors"
	"fmt"
	"lld/use_cases/snake_and_ladder/cell"
	snakeorladder "lld/use_cases/snake_and_ladder/snake_or_ladder"
)

var (
	ErrInvalidPosition = errors.New("invalid position")
)

type Board struct {
	cells []*cell.Cell
	size  int
}

func NewBoard(size int, snakes, ladders [][]int) (*Board, error) {
	board := &Board{
		size:  size,
		cells: make([]*cell.Cell, size),
	}

	for i := 0; i < size; i++ {
		board.cells[i] = cell.NewCell(i + 1)
	}

	for _, snake := range snakes {
		if err := board.createSnake(snake); err != nil {
			return nil, fmt.Errorf("create snake failed: %w", err)
		}
	}

	for _, ladder := range ladders {
		if err := board.createLadder(ladder); err != nil {
			return nil, fmt.Errorf("create ladder failed: %w", err)
		}
	}

	return board, nil
}

func (b *Board) createSnake(snakeRawInput []int) error {
	if len(snakeRawInput) != 2 {
		return errors.New("invalid snake entry: length should be 2")
	}

	from := snakeRawInput[0]
	to := snakeRawInput[1]

	if !b.IsValidPosition(from) || !b.IsValidPosition(to) {
		return fmt.Errorf("%s: %d, %d", ErrInvalidPosition, from, to)
	}

	snake, err := snakeorladder.NewSnake(from, to)
	if err != nil {
		return err
	}

	c, err := b.GetCell(from)
	if err != nil {
		return err
	}

	if err := c.SetSnakeOrLadder(snake); err != nil {
		return err
	}
	return nil
}

func (b *Board) createLadder(ladderRawInput []int) error {
	if len(ladderRawInput) != 2 {
		return errors.New("invalid ladder entry: length should be 2")
	}

	from := ladderRawInput[0]
	to := ladderRawInput[1]

	if !b.IsValidPosition(from) || !b.IsValidPosition(to) {
		return ErrInvalidPosition
	}

	ladder, err := snakeorladder.NewLadder(from, to)
	if err != nil {
		return err
	}

	c, err := b.GetCell(from)
	if err != nil {
		return err
	}

	if err := c.SetSnakeOrLadder(ladder); err != nil {
		return err
	}
	return nil
}

func (b *Board) IsValidPosition(pos int) bool {
	return 1 <= pos && pos <= len(b.cells)
}

func (b *Board) GetCell(pos int) (*cell.Cell, error) {
	if !b.IsValidPosition(pos) {
		return nil, ErrInvalidPosition
	}
	return b.cells[pos-1], nil
}

// Move returns next position after roll, if anything moved, is player won
func (b *Board) Move(currentPosition, roll int) (int, bool, bool) {
	nextPosition := currentPosition + roll
	if !b.IsValidPosition(nextPosition) {
		return currentPosition, false, false
	}
	for {
		c, _ := b.GetCell(nextPosition)
		if !c.HasSnakeOrLadder() {
			break
		}
		nextPosition = c.GetTo()
	}

	return nextPosition, true, nextPosition == len(b.cells)
}
