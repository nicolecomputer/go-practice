package aoc

import "strconv"

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
		} else {
			lastNum = num
		}
	}

	return int(firstNum)*10 + int(lastNum)
}
