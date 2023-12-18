package day7

import (
	"advent-of-code/common"
	"fmt"
	"regexp"
	"sort"
	"strconv"
)

type hand struct {
	cards    []string
	bid      int
	handType int
}

func convertCardToValue(card string) int {
	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 11
	case "T":
		return 10
	default:
		i, _ := strconv.Atoi(card)
		return i
	}
}

func SolveOne(file string) int {
	lines := common.GetFileLines(file)
	hands := getHands(lines)
	for _, hand := range hands {
		handType := findHandType(hand.cards)
		hand.handType = handType
	}
	orderHands(hands)
	printHands(hands)
	return calculateTotalWinnings(hands)
}

func getHands(lines []string) []*hand {
	var hands []*hand

	lineRegex := regexp.MustCompile(`(\w{5})\s(\d+)`)
	handRegex := regexp.MustCompile(`\w`)
	for _, line := range lines {
		match := lineRegex.FindStringSubmatch(line)
		bid, _ := strconv.Atoi(match[2])
		handMatch := handRegex.FindAllString(match[1], -1)
		hands = append(hands, &hand{cards: handMatch, bid: bid})
	}
	return hands
}

func findHandType(hand []string) int {
	unduplicatedHand := common.RemoveDuplicates(hand)
	instances := common.CountInstances(hand)

	// Five of a kind
	if len(unduplicatedHand) == 1 {
		return 7
		// Four of a kind or full house
	} else if len(unduplicatedHand) == 2 {
		// Four of a kind
		if common.MapContains(instances, 4) && common.MapContains(instances, 1) {
			return 6
			// Full house
		} else if common.MapContains(instances, 3) && common.MapContains(instances, 2) {
			return 5
		}
		// Three of a kind or two pair
	} else if len(unduplicatedHand) == 3 {
		// Three of a kind
		if common.MapContains(instances, 3) {
			return 4
			// Two pair
		} else if common.MapContains(instances, 2) {
			return 3
		}
		// One pair
	} else if len(unduplicatedHand) == 4 {
		return 2
	}
	// High card
	return 1
}

func orderHands(hands []*hand) {
	// Sort the hands by their hand type
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType != hands[j].handType {
			return hands[i].handType < hands[j].handType
		}
		for k := range hands[i].cards {
			if hands[i].cards[k] != hands[j].cards[k] {
				return convertCardToValue(hands[i].cards[k]) < convertCardToValue(hands[j].cards[k])
			}
		}
		return false
	})
}

func calculateTotalWinnings(hands []*hand) int {
	var total int
	for i, h := range hands {
		total += h.bid * (i + 1)
	}
	return total
}

func printHands(hands []*hand) {
	for _, h := range hands {
		fmt.Printf("%+v\n", h)
	}
}
