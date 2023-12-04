package day3

import (
	"fmt"
	"strings"

	"github.com/t00mas/aoc2023-go/parsers"
)

type Solution struct {
	Input   string
	Options *ProcessingOptions
}

type ProcessingOptions struct {
}

func (s *Solution) Init() {
	s.Input = input
}

func (s *Solution) Solve(star int) int {
	lines := parsers.ParseToLines(s.Input)
	// pad first/last line with dots
	dotLine := strings.Repeat(".", len(lines[0]))
	paddedLines := []string{}
	paddedLines = append(paddedLines, dotLine)
	paddedLines = append(paddedLines, lines...)
	paddedLines = append(paddedLines, dotLine)
	if star == 1 {
		return SumIfPartNumber(paddedLines)
	}
	if star == 2 {
		return SumGearRations(paddedLines)
	}
	return 0
}

func SumIfPartNumber(paddedLines []string) int {
	accSum := 0
	lineCursor := 1
	for {
		if lineCursor == len(paddedLines) {
			break
		}
		// pad line before/after
		currentLine := fmt.Sprintf(".%s.", paddedLines[lineCursor])
		numbers := getNumbers(currentLine, lineCursor)
		for _, n := range numbers {
			n = flagPartNumbers(n, paddedLines)
			if n.IsPartNumber {
				accSum = accSum + n.Value
			}
		}
		lineCursor++
	}

	return accSum
}

func SumGearRations(paddedLines []string) int {
	accGearRatios := 0
	lineCursor := 1
	for {
		if lineCursor == len(paddedLines) {
			break
		}
		currentLine := paddedLines[lineCursor]
		for i := 1; i < len(currentLine); i++ {
			if string(currentLine[i]) == "*" {
				topMidBottomLines := []string{}
				top := fmt.Sprintf(".%s.", paddedLines[lineCursor-1])
				topMidBottomLines = append(topMidBottomLines, top)
				mid := fmt.Sprintf(".%s.", paddedLines[lineCursor])
				topMidBottomLines = append(topMidBottomLines, mid)
				btm := fmt.Sprintf(".%s.", paddedLines[lineCursor+1])
				topMidBottomLines = append(topMidBottomLines, btm)

				nums := findNumbersInAround(i, topMidBottomLines)
				if len(nums) == 2 {
					accGearRatios = accGearRatios + (nums[0] * nums[1])
				}
			}
		}
		lineCursor++
	}
	return accGearRatios
}
