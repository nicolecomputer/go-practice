package main

import (
	"fmt"
	"os"
	"strings"

	"aoc/aoc"
)

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	str := string(bytes)
	str = strings.TrimSpace(str)

	part1 := aoc.TotalCalibrationValue(str)

	fmt.Println("Part 1:", part1)
}
