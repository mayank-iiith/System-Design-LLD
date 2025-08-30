package snakeorladder

import "errors"

type Snake struct {
	from int
	to   int
}

func NewSnake(from, to int) (SnakeOrLadder, error) {
	if to >= from {
		return nil, errors.New("invalid snake: tail >= head")

	}
	return &Snake{from, to}, nil
}

// GetTo implements SnakeOrLadder.
func (s *Snake) GetTo() int {
	return s.to
}

// GetType implements SnakeOrLadder.
func (s *Snake) GetType() string {
	return "Snake"
}
