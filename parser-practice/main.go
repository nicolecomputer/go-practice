package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/participle/v2"
)

type INI struct {
	Properties []*Property `@@*`
	Sections   []*Section  `@@*`
}

type Section struct {
	Identifier string      `"[" @Ident "]"`
	Properties []*Property `@@*`
}

type Property struct {
	Key   string `@Ident "="`
	Value *Value `@@`
}

type Value struct {
	String *string  `  @String`
	Float  *float64 `| @Float`
	Int    *int     `| @Int`
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

func main() {
	if len(os.Args) != 2 {
		fmt.Println("pass input as first arg")
		os.Exit(1)
	}

	inputFilename := os.Args[1]
	input := sanitizedInput(inputFilename)

	parser, err := participle.Build[INI]()
	if err != nil {
		fmt.Println("Error building grammar")
		os.Exit(1)
	}

	ast, err := parser.ParseString("", input)
	if err != nil {
		fmt.Println("Error parsing ", inputFilename)
		os.Exit(1)
	}

	fmt.Println(*ast.Properties[0].Value.Int * 10)
	s, _ := json.MarshalIndent(ast, "", "  ")

	fmt.Println(string(s))
}
