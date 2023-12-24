package day9

import (
	"advent-of-code/common"
	"slices"
	"strings"
)

func SolveOne(file string) int {
	lines := common.GetFileLines(file)
	var sequences [][]int
	for _, line := range lines {
		s := strings.Split(line, " ")
		sequences = append(sequences, common.ConvertToIntSlice(s))
	}

	var total int
	for _, sequence := range sequences {
		total += extrapolateForwards(sequence)
	}
	return total
}

func SolveTwo(file string) int {
	lines := common.GetFileLines(file)
	var sequences [][]int
	for _, line := range lines {
		s := strings.Split(line, " ")
		sequences = append(sequences, common.ConvertToIntSlice(s))
	}

	var total int
	for _, sequence := range sequences {
		total += extrapolateBackwards(sequence)
	}
	return total
}

// The next value is the sum of last item in the slices of differences
func extrapolateForwards(sequence []int) int {
	differences := findDifferences(sequence)

	var nextValue int
	for _, difference := range differences {
		nextValue += difference[len(difference)-1]
	}
	return nextValue + sequence[len(sequence)-1]
}

func findDifferences(sequence []int) [][]int {
	var differences [][]int
	currentSequence := sequence
	for !isBaseDifference(currentSequence) {
		nextSequence := findDifference(currentSequence)
		differences = append(differences, nextSequence)
		currentSequence = nextSequence
	}
	return differences
}

func isBaseDifference(difference []int) bool {
	return len(common.RemoveDuplicates(difference)) == 1
}

func findDifference(numbers []int) []int {
	var differences []int
	for i := 0; i < len(numbers)-1; i++ {
		differences = append(differences, numbers[i+1]-numbers[i])
	}
	return differences
}

func extrapolateBackwards(sequence []int) int {
	differences := findDifferences(sequence)
	slices.Reverse(differences)
	prevValue := differences[0][0]
	for _, difference := range differences[1:] {
		prevValue = difference[0] - prevValue
	}
	return sequence[0] - prevValue
}
