package main

import (
	"fmt"
	"os"
	"strings"

	"aoc/day01"
)

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

	input := sanitizedInput(os.Args[1])

	part1 := day01.TotalCalibrationValue(input, false)
	fmt.Println("Part 1:", part1)

	part2 := day01.TotalCalibrationValue(input, true)
	fmt.Println("Part 2:", part2)
}
