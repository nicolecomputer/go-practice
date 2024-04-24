package aoc

import (
	"strconv"
	"strings"
)

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

func lines(str string) []string {
	return strings.Split(str, "\n")
}

func TotalCalibrationValue(input string) int {
	total := 0
	return total
}
