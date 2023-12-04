package day2_test

import (
	"reflect"
	"testing"

	"github.com/t00mas/aoc2023-go/solutions/day2"
)

type ParseLineTestcase struct {
	Input              string
	Star               int
	ExpectedGameId     int
	ExpectedGameRounds []*day2.GameRound
}

func TestParseLine(t *testing.T) {
	tcs := []ParseLineTestcase{
		{
			Input:          "Game 100: 5 blue, 5 green; 7 blue, 15 green;",
			Star:           1,
			ExpectedGameId: 100,
			ExpectedGameRounds: []*day2.GameRound{
				{
					Green: 5,
					Blue:  5,
				},
				{
					Green: 15,
					Blue:  7,
				},
			},
		},
	}
	for _, tc := range tcs {
		gotGameId, gotGameRounds := day2.ParseLine(tc.Input)
		if tc.ExpectedGameId != gotGameId {
			t.Errorf("expected %d got %d", tc.ExpectedGameId, gotGameId)
		}
		if !reflect.DeepEqual(tc.ExpectedGameRounds, gotGameRounds) {
			t.Errorf("expected %+v got %+v", tc.ExpectedGameRounds, gotGameRounds)
		}
	}
}
