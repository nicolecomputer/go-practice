package day02

import (
	"fmt"
	"os"

	"github.com/alecthomas/participle/v2"
)

type GameRecord struct {
	Games []*Game `@@*`
}

type Game struct {
	Id          *int           `"Game" @Int ":"`
	RevealedSet []*RevealedSet `@@ ( ";" @@ )*`
}

type RevealedSet struct {
	Seen []*CubesSeen `@@ ( "," @@)*`
}

type CubesSeen struct {
	Count *int   "@Int"
	Color string "@Ident"
}

func ParseGame(input string) *Game {
	parser, err := participle.Build[Game]()
	if err != nil {
		fmt.Println("Error building grammar")
		os.Exit(1)
	}

	ast, err := parser.ParseString("", input)
	if err != nil {
		fmt.Println("Error parsing ")
		os.Exit(1)
	}

	return ast
}

func ParseGameRecord(input string) *GameRecord {
	parser, err := participle.Build[GameRecord]()
	if err != nil {
		fmt.Println("Error building grammar")
		os.Exit(1)
	}

	ast, err := parser.ParseString("", input)
	if err != nil {
		fmt.Println("Error parsing ")
		os.Exit(1)
	}

	return ast
}
