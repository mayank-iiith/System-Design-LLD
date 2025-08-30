package main

import (
	"lld/use_cases/snake_and_ladder/game"
	"log"
)

func main() {
	g, err := createGame()
	if err != nil {
		log.Fatalf("create game failed: %+v", err)
	}

	for !g.IsOver() {
		if err = g.RollDiceAndMove(); err != nil {
			log.Fatal(err)
		}
	}
}

func createGame() (*game.Game, error) {
	countOfCells := 100
	snakes := [][]int{
		{62, 5},
		{33, 6},
		{49, 9},
		{88, 16},
		{41, 20},
		{56, 53},
		{98, 64},
		{93, 73},
		{95, 75},
	}

	ladders := [][]int{
		{2, 37},
		{27, 46},
		{10, 32},
		{51, 68},
		{61, 79},
		{65, 84},
		{71, 91},
		{81, 100},
	}

	playerNames := []string{
		"Mayank",
		"Gaurav",
		"Sagar",
	}
	return game.NewGame(countOfCells, snakes, ladders, playerNames)
}
