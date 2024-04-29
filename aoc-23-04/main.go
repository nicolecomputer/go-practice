package main

import (
	"fmt"
	"os"

	"aoc/aoc"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("pass input as first arg")
		os.Exit(1)
	}

	inputFilename := os.Args[1]
	input := aoc.SanitizedInput(inputFilename)

	games := aoc.ParseGames(input)

	score := aoc.TotalScore(games)
	fmt.Println("Part 1:", score)
}
