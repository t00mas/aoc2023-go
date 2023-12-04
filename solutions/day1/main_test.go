package day1_test

import (
	"testing"

	"github.com/t00mas/aoc2023-go/solutions/day1"
)

type CalibrateLineTestcase struct {
	Input    string
	Star     int
	Expected int
}

func TestLine(t *testing.T) {
	tcs := []CalibrateLineTestcase{
		{
			Input:    "1abc2",
			Star:     1,
			Expected: 12,
		},
		{
			Input:    "pqr3stu8vwx",
			Expected: 38,
		},
		{
			Input:    "a1b2c3d4e5f",
			Expected: 15,
		},
		{
			Input:    "treb7uchet",
			Expected: 77,
		},
		{
			Input:    "two1nine",
			Star:     1,
			Expected: 11,
		},
		{
			Input:    "two1nine",
			Star:     2,
			Expected: 29,
		},
		{
			Input:    "eightwothree",
			Star:     2,
			Expected: 83,
		},
		{
			Input:    "abcone2threexyz",
			Star:     2,
			Expected: 13,
		},
		{
			Input:    "xtwone3four",
			Star:     2,
			Expected: 24,
		},
		{
			Input:    "4nineeightseven2",
			Star:     2,
			Expected: 42,
		},
		{
			Input:    "zoneight234",
			Star:     2,
			Expected: 14,
		},
		{
			Input:    "7pqrstsixteen",
			Star:     2,
			Expected: 76,
		},
	}
	for _, tc := range tcs {
		options := &day1.ProcessingOptions{ProcessNumWords: tc.Star == 2}
		got := day1.ProcessLine(tc.Input, options)
		if tc.Expected != got {
			t.Errorf("input: %s expected: %d got: %d", tc.Input, tc.Expected, got)
		}
	}
}

type ProcessTestCase struct {
	Input    string
	Star     int
	Expected int
}

func TestSmallInput(t *testing.T) {
	tcs := []ProcessTestCase{
		{
			Input:    "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet\n",
			Star:     1,
			Expected: 142,
		},
		{
			Input:    "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen\n",
			Star:     2,
			Expected: 281,
		},
	}
	for _, tc := range tcs {
		s := &day1.Solution{Input: tc.Input}
		got := s.Solve(tc.Star)
		if tc.Expected != got {
			t.Errorf("expected: %d got: %d", tc.Expected, got)
		}
	}
}
