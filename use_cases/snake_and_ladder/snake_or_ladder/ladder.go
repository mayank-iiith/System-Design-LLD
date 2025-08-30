package snakeorladder

import "errors"

type Ladder struct {
	from int
	to   int
}

func NewLadder(from, to int) (SnakeOrLadder, error) {
	if to <= from {
		return nil, errors.New("invalid entry: to should be greater than from")
	}

	return &Ladder{
		from: from,
		to:   to,
	}, nil
}

// GetTo implements SnakeOrLadder.
func (s *Ladder) GetTo() int {
	return s.to
}

// GetType implements SnakeOrLadder.
func (s *Ladder) GetType() string {
	return "Ladder"
}
