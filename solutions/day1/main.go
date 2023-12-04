package day1

import (
	"github.com/t00mas/aoc2023-go/parsers"
)

type Solution struct {
	Input   string
	Options *ProcessingOptions
}

type ProcessingOptions struct {
	ProcessNumWords bool
}

func (s *Solution) Init() {
	s.Input = input
}

func (s *Solution) Solve(star int) int {
	s.Options = &ProcessingOptions{ProcessNumWords: star == 2}
	calibration := 0
	parsed := parsers.ParseToLines(s.Input)
	for i := 0; i < len(parsed); i++ {
		if len(parsed[i]) < 1 {
			continue
		}
		calibration += ProcessLine(parsed[i], s.Options)
	}
	return calibration
}
