package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

func ProcessLineForSpelledNumbers(line string) string {
	result := strings.Clone(line)

	for i := 0; i < len(result); i++ {
		subStr := string(result[i:])

		if strings.HasPrefix(subStr, "two") {
			result = result[:i] + strings.Replace(subStr, "two", "2", 1)
		}

		if strings.HasPrefix(subStr, "nine") {
			result = result[:i] + strings.Replace(subStr, "nine", "9", 1)
		}

	}

	fmt.Println("result", result)
	return result
}

func CalibrationValue(line string) int {
	var firstNum int64 = -1
	var lastNum int64 = -1

	for i := 0; i < len(line); i++ {
		num, err := strconv.ParseInt(string(line[i]), 10, 64)
		if err != nil {
			continue
		}

		if firstNum == -1 {
			firstNum = num
		}

		lastNum = num
	}

	return int(firstNum)*10 + int(lastNum)
}

func allLines(str string) []string {
	return strings.Split(str, "\n")
}

// TODO:
// - Take CalibrationValue function as an argument
func TotalCalibrationValue(input string) int {
	total := 0
	lines := allLines(input)

	for _, line := range lines {
		total += CalibrationValue(line)
	}

	return total
}
