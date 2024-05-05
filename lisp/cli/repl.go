package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"lisp/lisp"
)

func welcomeMsg() string {
	return fmt.Sprintf("Go toy-lisp\ntype .exit to exit\n\n")
}

func prompt() string {
	return fmt.Sprintf("> ")
}

func evaluateStr(str string) {
	err, parsed := lisp.Parse(str)
	if err != nil {
		fmt.Fprintf(os.Stdout, "\033[0;31m %s\033[0m", err)
	}
	fmt.Println(parsed)
	fmt.Println()
}

func main() {
	fmt.Println(welcomeMsg())
	reader := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt())
	for reader.Scan() {
		text := reader.Text()
		if strings.EqualFold(".exit", text) {
			return
		} else {
			evaluateStr(text)
		}
		fmt.Print(prompt())
	}
	// Print an additional line if we encountered an EOF character
	fmt.Println()
}
