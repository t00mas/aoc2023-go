package day4_test

import (
	"testing"

	"github.com/t00mas/aoc2023-go/solutions/day4"
)

type TestCase struct {
	Input    string
	Star     int
	Expected int
}

func TestSolve(t *testing.T) {
	tcs := []TestCase{
		{
			Input:    "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11\n",
			Star:     1,
			Expected: 13,
		},
		{
			Input:    "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\nCard 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\nCard 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\nCard 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\nCard 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\nCard 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11\n",
			Star:     2,
			Expected: 30,
		},
	}
	for _, tc := range tcs {
		s := day4.Solution{
			Input: tc.Input,
		}
		got := s.Solve(tc.Star)
		if tc.Expected != got {
			t.Errorf("expected %+v got %+v", tc.Expected, got)
		}
	}
}
