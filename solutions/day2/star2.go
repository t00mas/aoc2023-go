package day2

func findAccRoundPower(gameLines []string) int {
	accRoundPower := 0

	for _, line := range gameLines {
		if len(line) == 0 {
			continue
		}
		_, gameRounds := ParseLine(line)
		minRound := findMinRound(gameRounds)
		accRoundPower += getRoundPower(minRound)
	}

	return accRoundPower
}

func findMinRound(rounds []*GameRound) *GameRound {
	minRound := &GameRound{}
	for _, round := range rounds {
		if minRound.Red < round.Red {
			minRound.Red = round.Red
		}
		if minRound.Green < round.Green {
			minRound.Green = round.Green
		}
		if minRound.Blue < round.Blue {
			minRound.Blue = round.Blue
		}
	}
	return minRound
}

func getRoundPower(round *GameRound) int {
	return round.Red * round.Green * round.Blue
}
