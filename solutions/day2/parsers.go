package day2

import (
	"fmt"
	"strconv"
	"strings"
)

type GameRound struct {
	Red   int
	Green int
	Blue  int
}

func ParseLine(line string) (int, []*GameRound) {
	game := strings.Split(line, ": ")
	fmt.Printf("game: %+v\n", game)
	gameHeader := strings.Split(game[0], " ")
	gameId, _ := strconv.Atoi(gameHeader[1])
	gameRounds := []*GameRound{}
	rawRounds := strings.Split(game[1], ";")
	for _, rawRound := range rawRounds {
		rawRound = strings.Trim(rawRound, " ")
		if len(rawRound) < 5 {
			continue
		}

		round := &GameRound{}
		values := strings.Split(rawRound, ", ")
		for _, value := range values {
			amountStr, color, _ := strings.Cut(value, " ")
			amount, _ := strconv.Atoi(string(amountStr))
			switch string(color) {
			case "red":
				round.Red = amount
			case "green":
				round.Green = amount
			case "blue":
				round.Blue = amount
			}
		}
		gameRounds = append(gameRounds, round)
	}

	return gameId, gameRounds
}
