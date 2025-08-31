package game

import (
	"errors"
	"fmt"
	"lld/use_cases/snake_and_ladder/board"
	"lld/use_cases/snake_and_ladder/dice"
	"lld/use_cases/snake_and_ladder/player"
)

type Game struct {
	dice    *dice.Dice
	board   *board.Board
	players []*player.Player
	turn    int
	over    bool
}

func NewGame(countOfCells int, snakes, ladders [][]int, playerNames []string) (*Game, error) {
	if len(playerNames) < 2 {
		return nil, errors.New("at least 2 players should be there")
	}

	d := dice.NewDice()

	b, err := board.NewBoard(countOfCells, snakes, ladders)
	if err != nil {
		return nil, err
	}

	players := make([]*player.Player, len(playerNames))
	for idx, playerName := range playerNames {
		players[idx] = player.NewPlayer(playerName)
	}

	return &Game{
		dice:    d,
		board:   b,
		players: players,
		turn:    0,
	}, nil
}

func (g *Game) PlayTurn() error {
	if g.IsOver() {
		return errors.New("game is already over")
	}
	p := g.players[g.turn]
	roll := g.dice.Roll()
	currentPosition := p.Position()
	newPosition, moved, win := g.board.Move(currentPosition, roll)
	if !moved {
		fmt.Printf("%s rolled a %d but cannot move from %d as it exceeds board limit, total gain: %d\n", p.Name(), roll, currentPosition, 0)
	} else {
		p.SetPosition(newPosition)
		fmt.Printf("%s rolled a %d and moved from %d to %d, total gain: %d\n", p.Name(), roll, currentPosition, newPosition, newPosition-currentPosition)
	}
	if win {
		fmt.Printf("%s has won the game!\n", p.Name())
		g.over = true
		return nil
	}

	g.setNextTurn()
	return nil
}

func (g *Game) setNextTurn() {
	g.turn = (g.turn + 1) % len(g.players)
}

func (g *Game) IsOver() bool {
	return g.over
}
