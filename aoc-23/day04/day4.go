package day04

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"

	"github.com/alecthomas/participle/v2"
)

type Game struct {
	Id             int   `"Card" @Int ":"`
	WinningNumbers []int `@Int+ "|"`
	NumbersOnCard  []int `@Int+`
}

func ParseGame(input string) *Game {
	parser, err := participle.Build[Game]()
	if err != nil {
		fmt.Println("Error building grammar")
		os.Exit(1)
	}

	game, err := parser.ParseString("", input)
	if err != nil {
		fmt.Println("Error parsing ")
		os.Exit(1)
	}

	return game
}

func ParseGames(input string) []*Game {
	var games []*Game
	for _, line := range strings.Split(input, "\n") {
		games = append(games, ParseGame(line))
	}

	return games
}

func Score(g *Game) int {
	var totalFound int = 0
	for _, winner := range g.WinningNumbers {
		if slices.Contains(g.NumbersOnCard, winner) {
			totalFound += 1
		}
	}

	if totalFound == 0 {
		return 0
	}

	return int(math.Pow(2, float64(totalFound-1)))
}

func TotalScore(games []*Game) int {
	score := 0
	for _, game := range games {
		score += Score(game)
	}
	return score
}
