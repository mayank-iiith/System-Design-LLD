package game

import (
	"errors"
	"lld/use_cases/snake_and_ladder/board"
	"lld/use_cases/snake_and_ladder/dice"
	"lld/use_cases/snake_and_ladder/player"
)

type Game struct {
	dice        *dice.Dice
	board       *board.Board
	players     []*player.Player
	currentTurn int
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
		players[idx] = player.NewPlayer(playerName, d, b)
	}

	return &Game{
		dice:        d,
		board:       b,
		players:     players,
		currentTurn: 0,
	}, nil
}

func (g *Game) RollDiceAndMove() error {
	if g.IsOver() {
		return errors.New("game is already over")
	}
	g.players[g.currentTurn].RollDiceAndMove()
	g.setNextTurn()
	return nil
}

func (g *Game) setNextTurn() {
	g.currentTurn = (g.currentTurn + 1) % len(g.players)
}

func (g *Game) IsOver() bool {
	return g.board.IsOver()
}
