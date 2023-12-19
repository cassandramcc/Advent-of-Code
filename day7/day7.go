package day7

import (
	"advent-of-code/common"
	"regexp"
	"slices"
	"sort"
	"strconv"
)

type hand struct {
	cards    []string
	bid      int
	handType int
}

func convertCardToValue(card string, joker bool) int {
	switch card {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		if joker {
			return 1
		}
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
		hand.handType = findHandType(hand.cards)
	}
	sortHands(hands, false)
	return calculateTotalWinnings(hands)
}

func SolveTwo(file string) int {
	lines := common.GetFileLines(file)
	hands := getHands(lines)
	for _, hand := range hands {
		hand.handType = findHandType(hand.cards)
		considerJokers(hand)
	}
	sortHands(hands, true)
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

func sortHands(hands []*hand, joker bool) {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType != hands[j].handType {
			return hands[i].handType < hands[j].handType
		}
		for k := range hands[i].cards {
			if hands[i].cards[k] != hands[j].cards[k] {
				return convertCardToValue(hands[i].cards[k], joker) < convertCardToValue(hands[j].cards[k], joker)
			}
		}
		return false
	})
}

// Must check that hand has jokers in it
func considerJokers(hand *hand) {
	if !slices.Contains(hand.cards, "J") {
		return
	}
	instances := common.CountInstances(hand.cards)
	// Four of a kind and full house becomes five of a kind
	if hand.handType == 6 || hand.handType == 5 {
		hand.handType = 7

		// three of a kind becomes four of a kind
	} else if hand.handType == 4 {
		hand.handType = 6

		// two pair
	} else if hand.handType == 3 {
		if instances["J"] == 2 {
			// becomes four of a kind
			hand.handType = 6

			// becomes full house
		} else if instances["J"] == 1 {
			hand.handType = 5
		}

		// one pair becomes three of a kind
	} else if hand.handType == 2 {
		hand.handType = 4

		// high card becomes one pair
	} else if hand.handType == 1 {
		hand.handType = 2
	}
}

func calculateTotalWinnings(hands []*hand) int {
	var total int
	for i, h := range hands {
		total += h.bid * (i + 1)
	}
	return total
}
