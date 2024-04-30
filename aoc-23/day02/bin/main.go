package main

import (
	"fmt"
	"os"

	"aoc/day02"
	"aoc/util"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("pass input as first arg")
		os.Exit(1)
	}

	inputFilename := os.Args[1]
	input := util.SanitizedInput(inputFilename)

	gameRecord := day02.ParseGameRecord(input)

	score := day02.TotalValueForPossibleGames(
		day02.CubeConfiguration{Red: 12, Blue: 13, Green: 14},
		gameRecord)
	fmt.Println("Part 1:", score)
}
