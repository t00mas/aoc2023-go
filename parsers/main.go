package parsers

import "strings"

func ParseToLines(inputStr string) []string {
	lines := strings.Split(inputStr, "\n")
	return lines[:len(lines)-1]
}
