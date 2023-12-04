package day3_test

import (
	"reflect"
	"testing"

	"github.com/t00mas/aoc2023-go/solutions/day3"
)

type FindAndFlagNumbersInLinesTestcases struct {
	InputLines   []string
	InputLineIdx int
	Expected     []*day3.Number
}

var originalExampleInput = "467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..\n"

type SolveTestcase struct {
	Input    string
	Star     int
	Expected int
}

func TestSolve(t *testing.T) {
	tcs := []SolveTestcase{
		{
			Input:    originalExampleInput,
			Star:     1,
			Expected: 4361,
		},
		{
			Input:    "..152*141.\n",
			Star:     1,
			Expected: 293,
		},
		{
			Input:    "*...\n.5..\n...7\n",
			Star:     1,
			Expected: 5,
		},
	}
	for _, tc := range tcs {
		s := day3.Solution{
			Input: tc.Input,
		}
		got := s.Solve(tc.Star)
		if !reflect.DeepEqual(tc.Expected, got) {
			t.Errorf("expected %+v got %+v", tc.Expected, got)
		}
	}
}
