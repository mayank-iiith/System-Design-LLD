package main

import (
	"lld/use_cases/snake_and_ladder/game"
	"log"
)

func main() {
	countOfCells, snakes, ladders, playerNames := readInput()
	g, err := game.NewGame(countOfCells, snakes, ladders, playerNames)
	if err != nil {
		log.Fatalf("create game failed: %+v", err)
	}

	for !g.IsOver() {
		if err = g.PlayTurn(); err != nil {
			log.Fatal(err)
		}
	}
}

func readInput() (int, [][]int, [][]int, []string) {
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
	return countOfCells, snakes, ladders, playerNames
}
