package day1

import (
	"advent-of-code/common"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func SolveOne(file string) int {
	lines := common.GetFileLines(file)

	digit := regexp.MustCompile(`\d`)
	var total int
	for _, line := range lines {
		digits := digit.FindAllString(line, -1)
		firstDigit := digits[0]
		lastDigit := digits[len(digits)-1]
		fullNumber, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			fmt.Println("bad conversion!")
			os.Exit(1)
		}
		total += fullNumber
	}
	return total
}

func convertDigitString(digit string) int {
	digitStrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	// Try to convert the digit to an int
	// If it fails, the digit is text based and needs to be converted via an index in the digitStrings slice
	i, err := strconv.Atoi(digit)
	if err != nil {
		i = slices.Index(digitStrings, digit) + 1
	}
	return i
}

func findDigits(line string) int {
	digitStrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	firstIndex := math.MaxInt
	lastIndex := -math.MaxInt

	var firstDigit string
	var lastDigit string

	// Find the first and last string based digit
	for _, digit := range digitStrings {
		firstIdx := strings.Index(line, digit)
		lastIdx := strings.LastIndex(line, digit)
		if firstIdx != -1 && firstIdx < firstIndex {
			firstIndex = firstIdx
			firstDigit = digit
		}

		if lastIdx != -1 && lastIdx > lastIndex {
			lastIndex = lastIdx
			lastDigit = digit
		}
	}

	firstInt := convertDigitString(firstDigit)
	lastInt := convertDigitString(lastDigit)

	return firstInt*10 + lastInt
}

func SolveTwo(file string) int {
	lines := common.GetFileLines(file)

	var total int

	for _, line := range lines {
		total += findDigits(line)
	}
	return total
}
