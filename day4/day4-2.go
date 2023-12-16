package day4

import (
	"advent-of-code/common"
)

func SolveTwo(file string) int {
	lines := common.GetFileLines(file)

	wins := calculateAllGameWins(lines)
	total := calculateNumCards(wins)

	return total
}

func calculateAllGameWins(lines []string) []int {
	var gameWins []int
	for _, line := range lines {
		winners, played := extractNumbers(line)
		wins := calculateGameWins(winners, played)
		gameWins = append(gameWins, wins)
	}
	return gameWins
}

func calculateNumCards(wins []int) int {
	var numCards int

	cards := make(map[int]int, len(wins))

	// Fill out the card instances map with 1 card initially
	for i := range wins {
		cards[i+1] = 1
	}

	// Calculate how many extra cards appear
	for gameID, win := range wins {
		for i := gameID + 1; i < gameID+win+1; i++ {
			cards[i+1] += cards[gameID+1]
		}
	}

	// Total up the card instances
	for _, card := range cards {
		numCards += card
	}

	return numCards
}
