/*:
# CracklePop!

Hello, Adventurer! 💖

Welcome to the secret headquarters. We have snacks! But, the snacks are only
available to those that know the secret code. A friendly colleague may come up
to you and say "3" and you'll have to tell them back "Crackle" then they **might** share their pretzels and hummus. The codewords are based on an
algorithm. But don't worry, you can use this handy program to generate a table of
all the code words so you'll always be ready for snack time. 🥨

[Video Demo]

*[In a more serious tone]* CracklePop is a program in the vein of
[FizzBuzz](https://en.wikipedia.org/wiki/Fizz_buzz). It aims to be the very best
FizzBuzz it can possibly be, though, with features to return a single value or a
table of values.

To get a complete table of values you can call `cracklepop` with no arguments. If you'd
like just one value you pass a number of the program `cracklepop 109`. There are flags
for setting start and end and dialect. Oh! Dialect lets you control what words are output.
This version includes `default` (Crackle, Pop, CracklePop) and `popcorn` (Pop, Corn, Popcorn).
These dialects provide a nice place to internationalize or personalize this program
to different use cases.
:*/

package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/*:
## Core Algorithm

This is that we're here for ✨ (it's the code you asked for).

The core algorithm checks divisibility by using the modulo operator
and returns a response. This is a pretty standard implementation. The only
thing that's out of the ordinary is that CodeWordResponse() doesn't output
directly to standard out. This is so the response can be formatted depending
on whether we're showing a single value or a table of values.
:*/

func CodeWordResponse(value int) Response {
	if value%3 == 0 && value%5 == 0 {
		return CracklePop{}
	} else if value%3 == 0 {
		return Crackle{}
	} else if value%5 == 0 {
		return Pop{}
	}

	return Number{Value: value}
}

/*:
## Support Code

These are the basic types and functions for creating code words.

The `CodeWords()` function will generate a run of numbers that's used
for generating tables.
:*/

func CodeWords(start int, end int) []CodeWord {
	result := []CodeWord{}

	for i := start; i <= end; i++ {
		result = append(result, CodeWord{Challenge: i, Response: CodeWordResponse(i)})
	}

	return result
}

type CodeWord struct {
	Challenge int
	Response  Response
}

type Response interface {
	Text(d Dialect) string
	FormattedText(d Dialect) string
}

type (
	Crackle    struct{}
	Pop        struct{}
	CracklePop struct{}
	Number     struct{ Value int }
)

/*:
## Dialects

This version of the problem uses the words Crackle, Pop, and CracklePop but
there are other versions of this problem. Some of them are fizzy. And what
about our friends who speak other languages? This version of the program
includes one additional dialect, `popcorn`.

If this program were to continue to grow this should be moved out into a
separate file or even maintained separately from the code. (Maybe we need
to submit strings to folks who speak the language and they usually don't
like code)
:*/

type Dialect struct {
	Greeting   string
	Crackle    string
	Pop        string
	CracklePop string
	Warning    string
}

func (d Dialect) longestWord() int {
	return max(len(d.Crackle), len(d.Pop), len(d.CracklePop))
}

var dialects map[string]Dialect = map[string]Dialect{
	"default": {
		Greeting:   "Hello, Adventurer. ✨",
		Crackle:    "Crackle",
		Pop:        "Pop",
		CracklePop: "CracklePop",
	},
	"popcorn": {
		Greeting:   "Hello, Friend. ✨",
		Crackle:    "Pop 🎊",
		Pop:        "Corn 🌽",
		CracklePop: "Popcorn 🍿",
	},
}

/*:
### Plain Text versions

This formats each Response into plain text with no formatting
:*/

func (c Crackle) Text(d Dialect) string {
	return d.Crackle
}

func (p Pop) Text(d Dialect) string {
	return d.Pop
}

func (cp CracklePop) Text(d Dialect) string {
	return d.CracklePop
}

func (n Number) Text(d Dialect) string {
	return fmt.Sprintf("%d", n.Value)
}

/*:
### Formatted Text

This formats each Response into rich text for a terminal with color

:*/

func (c Crackle) FormattedText(d Dialect) string {
	return colorize(red, rpad(c.Text(d), d.longestWord()))
}

func (p Pop) FormattedText(d Dialect) string {
	return colorize(blue, rpad(p.Text(d), d.longestWord()))
}

func (cp CracklePop) FormattedText(d Dialect) string {
	return colorize(yellow, rpad(cp.Text(d), d.longestWord()))
}

func (n Number) FormattedText(d Dialect) string {
	return rpad(n.Text(d), d.longestWord())
}

/*:
## Terminal Code

This code deals with the terminal and has functions for learning about the
environment and for setting things like colors
:*/

func colorize(format string, s string) string {
	return fmt.Sprintf("%s%s%s%s", "\033[", format, s, "\033[0m")
}

const (
	red    = "0;31m"
	blue   = "0;34m"
	yellow = "0;33m"
	white  = "0;37m"
	gray   = "0;90m"
)

func terminalColumns() int {
	const defaultColumnSize int = 80

	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return defaultColumnSize
	}

	_, sizeAsString, _ := strings.Cut(string(out[:]), " ")
	size, err := strconv.Atoi(strings.TrimSpace(sizeAsString))
	if err != nil {
		return defaultColumnSize
	}

	return size
}

/*:
# Utility Functions
:*/
// %-*s pads a string to a length with the padding at the end
func rpad(s string, length int) string {
	return fmt.Sprintf("%-*s", length, s)
}

func fatalError(explanation string) {
	error := colorize(red, explanation)
	os.Stderr.WriteString(error)
	os.Stderr.WriteString("\n")
	os.Exit(1)
}

/*:
## Formatting
:*/

func formattedCodeWord(codeword CodeWord, dialect Dialect) string {
	return fmt.Sprintf(" %03d: %s", codeword.Challenge, codeword.Response.FormattedText(dialect))
}

func totalLengthForEntry(dialect Dialect) int {
	// <space><3 digits for numbers><colon><space><lengthOfLongestWord>
	return 1 + 3 + 1 + 1 + dialect.longestWord()
}

// Behold- the cutest fizzbuzz

/*
Only generates codes in the range 0-999
*/
func main() {
	// Check if there is exactly 1 argument

	dialectName := flag.String("dialect", "default", "the dialect to generate code words in")
	start := flag.Int("start", 1, "the starting value to generate codes from")
	end := flag.Int("end", -1, "the last value. If -1 is passed then start+99 will be the end")

	flag.Parse()

	if *end == -1 {
		*end = min(*start+99, 999)
	}

	if *start < 0 {
		fatalError("start must be greater than 0")
	}
	if *end > 999 {
		fatalError("end must be less than 1000")
	}
	if *start > *end {
		fatalError("start cannot be greater than end")
	}

	dialect, dialectPresent := dialects[*dialectName]
	if !dialectPresent {
		fatalError(fmt.Sprintf("Could not find dialect %s", *dialectName))
	}

	if len(flag.Args()) == 1 {
		value, err := strconv.Atoi(flag.Arg(0))
		if err != nil {
			fatalError("generating codewords only works with whole numbers")
		}
		res := CodeWordResponse(value).Text(dialect)
		fmt.Println(res)
	} else if len(flag.Args()) > 1 {
		fatalError("Unknown arguments. Progam takes 1 arugment to generate codeword")
	} else {

		// Generate code words
		codes := CodeWords(*start, *end)

		// This leaves some space on the table because it adds 2 to every entry for a left and right divider
		// and most entries share a left and right divider. Since it never overflows this is good enough
		columns := (terminalColumns()) / (totalLengthForEntry(dialect) + 2)

		fmt.Println(colorize("1m", dialect.Greeting))
		fmt.Println("Here are today's code words")

		// Length for each entry * columns - the last column doesn't have a join character
		dataLength := totalLengthForEntry(dialect)*columns + columns - 1
		fmt.Printf("%s%s%s", "╭", strings.Repeat("─", dataLength), "╮\n")

		// This needs A LOT of explaining
		for row := 0; row < len(codes); row += columns {
			result := []string{}
			for col := 0; col < columns; col++ {
				pos := row + col
				if pos < len(codes) {
					result = append(result, formattedCodeWord(codes[pos], dialect))
				} else {
					// This adds a space so that columns always line up
					result = append(result, strings.Repeat(" ", totalLengthForEntry(dialect)))
				}
			}
			fmt.Printf("│%s│\n", strings.Join(result, "║"))
		}

		fmt.Printf("%s%s%s", "╰", strings.Repeat("─", dataLength), "╯\n")
		fmt.Println(colorize(gray, "If you don't give the right codeword we won't share our snacks. 🥨 🍿 🍪"))
	}
}

/*:
## Things I learned

- Go string formatting
- Terminal box characters
- go flags

## Frustrations
- Go doesn't have sum types so enums are weakly typed and checking values isn't exhausitve
:*/
