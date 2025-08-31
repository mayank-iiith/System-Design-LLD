package snakeorladder

import "errors"

type Ladder struct {
	from int
	to   int
}

func NewLadder(from, to int) (SnakeOrLadder, error) {
	if to <= from {
		return nil, errors.New("invalid ladder: top <= base")
	}

	return &Ladder{from, to}, nil
}

// GetTo implements SnakeOrLadder.
func (l *Ladder) GetTo() int {
	return l.to
}

// GetType implements SnakeOrLadder.
func (*Ladder) GetType() string {
	return "Ladder"
}
