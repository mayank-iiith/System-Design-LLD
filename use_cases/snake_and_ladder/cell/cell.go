package cell

import (
	"fmt"
	snakeorladder "lld/use_cases/snake_and_ladder/snake_or_ladder"
)

type Cell struct {
	num           int
	snakeOrLadder snakeorladder.SnakeOrLadder
}

func NewCell(num int) *Cell {
	return &Cell{
		num: num,
	}
}

func (c *Cell) GetNum() int {
	return c.num
}

func (c *Cell) HasSnakeOrLadder() bool {
	return c.snakeOrLadder != nil
}

func (c *Cell) SetSnakeOrLadder(snakeOrLadder snakeorladder.SnakeOrLadder) error {
	if c.HasSnakeOrLadder() {
		return fmt.Errorf("already has %s at cell %d", c.snakeOrLadder.GetType(), c.GetNum())
	}
	c.snakeOrLadder = snakeOrLadder
	return nil
}

func (c *Cell) GetTo() int {
	if c.HasSnakeOrLadder() {
		return c.snakeOrLadder.GetTo()
	}
	return c.num
}
