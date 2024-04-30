package day01

import (
	"strconv"
	"strings"
)

func ProcessLineForSpelledNumbers(line string) string {
	result := ""

	for i := 0; i < len(line); i++ {
		_, err := strconv.Atoi(string(line[i]))
		if err == nil {
			result = result + string(line[i])
		}

		subStr := string(line[i:])

		if strings.HasPrefix(subStr, "one") {
			result = result + "1"
		} else if strings.HasPrefix(subStr, "two") {
			result = result + "2"
		} else if strings.HasPrefix(subStr, "three") {
			result = result + "3"
		} else if strings.HasPrefix(subStr, "four") {
			result = result + "4"
		} else if strings.HasPrefix(subStr, "five") {
			result = result + "5"
		} else if strings.HasPrefix(subStr, "six") {
			result = result + "6"
		} else if strings.HasPrefix(subStr, "seven") {
			result = result + "7"
		} else if strings.HasPrefix(subStr, "eight") {
			result = result + "8"
		} else if strings.HasPrefix(subStr, "nine") {
			result = result + "9"
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
