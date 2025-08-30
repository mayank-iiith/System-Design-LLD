package dice

import (
	"math/rand/v2"
)

type Dice struct {
	min, max int
}

func NewDice() *Dice {
	return &Dice{
		min: 1,
		max: 6,
	}
}

func (d *Dice) Roll() int {
	return rand.IntN(d.max-d.min+1) + d.min
}
