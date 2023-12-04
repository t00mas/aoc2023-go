package day2

func findAccGameIdSum(gameLines []string, opts *ProcessingOptions) int {
	accIdSum := 0

	for _, line := range gameLines {
		if len(line) == 0 {
			continue
		}
		gameId, gameRounds := ParseLine(line)
		if !isGameValid(gameRounds, opts.MaxRoundValues) {
			continue
		}
		accIdSum += gameId
	}
	return accIdSum
}

func isRoundValid(round, maxValues *GameRound) bool {
	if round.Red > maxValues.Red {
		return false
	}
	if round.Green > maxValues.Green {
		return false
	}
	if round.Blue > maxValues.Blue {
		return false
	}
	return true
}

func isGameValid(rounds []*GameRound, maxValues *GameRound) bool {
	for _, round := range rounds {
		if !isRoundValid(round, maxValues) {
			return false
		}
	}
	return true
}
