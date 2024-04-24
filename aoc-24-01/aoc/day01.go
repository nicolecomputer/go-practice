package aoc

import (
	"strconv"
	"strings"
)

func ProcessLineForSpelledNumbers(line string) string {
	result := strings.Clone(line)

	for i := 0; i < len(result); i++ {
		subStr := string(result[i:])

		if strings.HasPrefix(subStr, "one") {
			result = result[:i] + strings.Replace(subStr, "one", "1", 1)
		} else if strings.HasPrefix(subStr, "two") {
			result = result[:i] + strings.Replace(subStr, "two", "2", 1)
		} else if strings.HasPrefix(subStr, "three") {
			result = result[:i] + strings.Replace(subStr, "three", "3", 1)
		} else if strings.HasPrefix(subStr, "four") {
			result = result[:i] + strings.Replace(subStr, "four", "4", 1)
		} else if strings.HasPrefix(subStr, "five") {
			result = result[:i] + strings.Replace(subStr, "five", "5", 1)
		} else if strings.HasPrefix(subStr, "six") {
			result = result[:i] + strings.Replace(subStr, "six", "6", 1)
		} else if strings.HasPrefix(subStr, "seven") {
			result = result[:i] + strings.Replace(subStr, "seven", "7", 1)
		} else if strings.HasPrefix(subStr, "eight") {
			result = result[:i] + strings.Replace(subStr, "eight", "8", 1)
		} else if strings.HasPrefix(subStr, "nine") {
			result = result[:i] + strings.Replace(subStr, "nine", "9", 1)
		}
	}

	return result
}

func CalibrationValue(line string, filterSpelledOutWords bool) int {
	if filterSpelledOutWords {
		line = ProcessLineForSpelledNumbers(line)
	}

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

func TotalCalibrationValue(input string, filterSpelledOutWords bool) int {
	total := 0
	lines := allLines(input)

	for _, line := range lines {
		total += CalibrationValue(line, filterSpelledOutWords)
	}

	return total
}
