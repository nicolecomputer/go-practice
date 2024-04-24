package main

import (
	"fmt"
	"os"
	"strings"

	"aoc/aoc"
)

/*
Next:
- Refactor main to take filename as arg
- Add a SanitizedInput string that reads from file, trims and gives back a string
*/
func main2() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(bytes)
	str = strings.TrimSpace(str)

	part1 := aoc.TotalCalibrationValue(str)

	fmt.Println("Part 1:", part1)
}

func main() {
	aoc.ProcessLineForSpelledNumbers("two1nine")
}
