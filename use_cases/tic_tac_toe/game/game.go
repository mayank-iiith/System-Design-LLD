package game

import (
	"errors"
	"fmt"
	"lld/use_cases/tic_tac_toe/grid"
	"lld/use_cases/tic_tac_toe/piece"
	"lld/use_cases/tic_tac_toe/player"
	"strconv"
	"strings"
)

type Game struct {
	grid    *grid.Grid
	players []*player.Player
	turn    int
	over    bool
}

func NewGame(dimension int) *Game {
	return &Game{
		grid:    grid.NewGrid(dimension),
		players: make([]*player.Player, 0),
	}
}

func (g *Game) SetPlayers(players [][]string) error {
	if len(players) != 2 {
		return errors.New("only 2 players are allowed")
	}
	for i, p := range players {
		if len(p) != 2 {
			return errors.New("invalid input")
		}

		pieceType, err := piece.GetPiece(p[0])
		if err != nil {
			return err
		}

		if pieceType == "X" {
			g.turn = i
		}

		g.players = append(g.players, player.NewPlayer(p[1], pieceType))
	}
	return nil
}

func (g *Game) SetNextTurn() {
	g.turn = (g.turn + 1) % len(g.players)
}

func (g *Game) PlayTurn(move string) error {
	if g.over {
		return errors.New("game is already over")
	}
	if move == "exit" {
		g.over = true
		return nil
	}
	splits := strings.Split(move, " ")
	if len(splits) != 2 {
		return errors.New("invalid move")
	}
	posX, err := strconv.Atoi(splits[0])
	if err != nil {
		return errors.New("invalid move")
	}

	posY, err := strconv.Atoi(splits[1])
	if err != nil {
		return errors.New("invalid move")
	}
	return g.playTurn(posX, posY)
}

func (g *Game) playTurn(posX, posY int) error {
	if !g.grid.IsValidPosition(posX, posY) || g.grid.HasPiece(posX, posY) {
		fmt.Printf("Invalid Move, Play again...\n\n")
		return nil
	}
	p := g.players[g.turn]
	over, win := g.grid.PlacePiece(p.PieceType(), posX, posY)
	g.grid.Print()
	g.over = over
	if win {
		fmt.Printf("%s won the game with pieceType %s\n", p.Name(), p.PieceType().String())
		return nil
	} else if over {
		fmt.Println("Game Over")
		return nil
	}
	g.SetNextTurn()
	return nil
}
