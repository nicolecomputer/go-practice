package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

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
	Seen *[]CubesSeen `@@ ( "," @@)*`
}

type CubesSeen struct {
	Count *int   "@Int"
	Color string "@Ident"
}

// Code to run the parser

func sanitizedInput(path string) string {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	str := string(bytes)
	str = strings.TrimSpace(str)

	return str
}

func parseGameRecord(input string) *GameRecord {
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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("pass input as first arg")
		os.Exit(1)
	}

	inputFilename := os.Args[1]
	input := sanitizedInput(inputFilename)
	gameRecord := parseGameRecord(input)

	s, _ := json.MarshalIndent(gameRecord, "", "  ")

	fmt.Println(string(s))
}
