package day4

import (
	"advent-of-code/common"
	"math"
	"regexp"
	"slices"
	"strconv"
)

func SolveOne(file string) int {
	lines := common.GetFileLines(file)

	var total int
	for _, line := range lines {
		winners, played := extractNumbers(line)
		wins := calculateGameWins(winners, played)
		total += calculateGameScore(wins)
	}

	return total
}

func extractNumbers(line string) ([]int, []int) {
	numbersRegex := regexp.MustCompile(`((?:\d+\s+)*)\|((?:\s+\d+)*)`)
	digitsRegex := regexp.MustCompile(`\d+`)
	matches := numbersRegex.FindStringSubmatch(line)
	winners := matches[1]
	played := matches[2]

	winningNumberStrings := digitsRegex.FindAllString(winners, -1)
	playedNumbersString := digitsRegex.FindAllString(played, -1)

	var winningNumbers []int
	var playedNumbers []int
	for _, w := range winningNumberStrings {
		i, _ := strconv.Atoi(w)
		winningNumbers = append(winningNumbers, i)
	}

	for _, p := range playedNumbersString {
		i, _ := strconv.Atoi(p)
		playedNumbers = append(playedNumbers, i)
	}

	return winningNumbers, playedNumbers
}

func calculateGameWins(winners []int, played []int) int {
	var wins int
	// I am assuming that the played cannot have duplicate numbers
	for _, play := range played {
		if slices.Contains(winners, play) {
			wins += 1
		}
	}
	return wins
}

func calculateGameScore(wins int) int {
	if wins == 0 {
		return 0
	}
	return int(math.Pow(2, float64(wins-1)))
}
