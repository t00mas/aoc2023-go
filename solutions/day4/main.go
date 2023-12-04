package day4

import (
	"github.com/t00mas/aoc2023-go/parsers"
)

type Solution struct {
	Input   string
	Star    int
	Options *ProcessingOptions
}

type ProcessingOptions struct {
}

func (s *Solution) Init() {
	s.Input = input
}

func (s *Solution) Solve(star int) int {
	lines := parsers.ParseToLines(s.Input)
	cards := []*Card{}
	for _, l := range lines {
		card := &Card{}
		card.Parse(l)
		cards = append(cards, card)
	}

	if star == 1 {
		totalPoints := 0
		for _, c := range cards {
			totalPoints = totalPoints + c.PointValue
		}
		return totalPoints
	}

	if star == 2 {
		cardCountMap := map[int]int{}
		for _, c := range cards {
			cardCountMap[c.Index] = cardCountMap[c.Index] + 1
			for i := 0; i < c.Matches; i++ {
				cardCountMap[c.Index+i+1] = cardCountMap[c.Index+i+1] + cardCountMap[c.Index]
			}
		}
		total := 0
		for _, n := range cardCountMap {
			total = total + n
		}
		return total
	}

	return 0
}
