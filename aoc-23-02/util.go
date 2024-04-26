package aoc2302

import (
	"fmt"
	"os"
	"strings"
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
