package main

import (
	"fmt"
	"lld/use_cases/tic_tac_toe/game"
	"log"
)

func main() {
	fmt.Printf("Tic Tac Toe\n\n")
	dimension, players, moves := readInput1()
	//dimension, players, moves := readInput2()
	//dimension, players, moves := readInput3()

	g := game.NewGame(dimension)
	err := g.SetPlayers(players)
	if err != nil {
		log.Fatal(err)
	}

	for _, move := range moves {
		if err := g.PlayTurn(move); err != nil {
			log.Fatal(err)
		}
	}
}

func readInput1() (int, [][]string, []string) {
	dimension := 3
	players := [][]string{
		{"X", "Gaurav"},
		{"O", "Sagar"},
	}

	moves := []string{
		"2 2",
		"1 3",
		"1 1",
		"1 2",
		"2 2",
		"3 3",
		"exit",
	}
	return dimension, players, moves
}

func readInput2() (int, [][]string, []string) {
	dimension := 3
	players := [][]string{
		{"X", "Gaurav"},
		{"O", "Sagar"},
	}

	moves := []string{
		"2 3",
		"1 2",
		"2 2",
		"2 1",
		"1 1",
		"3 3",
		"3 2",
		"3 1",
		"1 3",
		"exit",
	}
	return dimension, players, moves
}

func readInput3() (int, [][]string, []string) {
	dimension := 3
	players := [][]string{
		{"X", "Gaurav"},
		{"O", "Sagar"},
	}

	moves := []string{
		"exit",
	}
	return dimension, players, moves
}
