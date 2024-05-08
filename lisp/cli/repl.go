package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	parser "lisp/parse"
)

func welcomeMsg() string {
	return fmt.Sprintf("Go toy-lisp\ntype .exit to exit\n")
}

func prompt() string {
	return fmt.Sprintf("\n> ")
}

func evaluateStr(str string) {
	err, parsed := parser.Parse(str)
	if err != nil {
		fmt.Fprintf(os.Stdout, "\033[0;31m %s\033[0m", err)
	}
	fmt.Println(parsed)
}

func main() {
	fmt.Println(welcomeMsg())
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt())
	for reader.Scan() {
		text := reader.Text()
		if strings.EqualFold(".exit", text) {
			return
		} else if strings.HasPrefix(text, ".lex") {
			input := strings.TrimPrefix(text, ".lex")
			var results []string = []string{}

			for _, t := range parser.Lex(input) {
				token := fmt.Sprintf("<%s>", t)
				results = append(results, token)
			}

			result := strings.Join(results, ", ")

			fmt.Println(result)
		} else if strings.HasPrefix(text, ".tokenize") {
			input := strings.TrimPrefix(text, ".tokenize")
			var results []string = []string{}

			for _, token := range parser.Tokenize(parser.Lex(input)) {
				results = append(results, token.String())
			}

			result := strings.Join(results, ", ")

			fmt.Println(result)
		} else {
			evaluateStr(text)
		}
		fmt.Print(prompt())
	}
	// Print an additional line if we encountered an EOF character
	fmt.Println()
}
