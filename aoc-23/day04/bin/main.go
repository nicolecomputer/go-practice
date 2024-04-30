package main

import (
	"fmt"
	"os"

	"aoc/day04"
	"aoc/util"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("pass input as first arg")
		os.Exit(1)
	}

	inputFilename := os.Args[1]
	input := util.SanitizedInput(inputFilename)

	games := day04.ParseGames(input)

	score := day04.TotalScore(games)
	fmt.Println("Part 1:", score)
}
