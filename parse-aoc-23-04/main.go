package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/participle/v2"
)

type Game struct {
	Id             int   `"Card" @Int ":"`
	WinningNumbers []int `@Int+ "|"`
	NumbersOnCard  []int `@Int+`
}

func sanitizedInput(path string) string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	str := string(bytes)
	str = strings.TrimSpace(str)

	return str
}

func parseGames(input string) []*Game {
	parser, err := participle.Build[Game]()
	if err != nil {
		fmt.Println("Error building grammar")
		os.Exit(1)
	}

	var games []*Game
	for _, line := range strings.Split(input, "\n") {
		game, err := parser.ParseString("", line)
		if err != nil {
			fmt.Println("Error parsing ")
			os.Exit(1)
		}
		games = append(games, game)
	}

	return games
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("pass input as first arg")
		os.Exit(1)
	}

	inputFilename := os.Args[1]
	input := sanitizedInput(inputFilename)

	g := parseGames(input)

	s, _ := json.MarshalIndent(g, "", "  ")
	fmt.Println(string(s))
}
