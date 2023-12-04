package day2

import "github.com/t00mas/aoc2023-go/parsers"

type Solution struct {
	Input   string
	Options *ProcessingOptions
}

type ProcessingOptions struct {
	MaxRoundValues *GameRound
}

func (s *Solution) Init() {
	s.Input = input
}

func (s *Solution) Solve(star int) int {
	gameLines := parsers.ParseToLines(s.Input)

	if star == 1 {
		opts := &ProcessingOptions{
			MaxRoundValues: &GameRound{
				Red:   12,
				Green: 13,
				Blue:  14,
			},
		}
		return findAccGameIdSum(gameLines, opts)
	}

	if star == 2 {
		return findAccRoundPower(gameLines)
	}

	return 0
}
