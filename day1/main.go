package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	inputs = "inputs/"
)

func main() {
	args := os.Args[1:]
	fmt.Println(solveTwo(args[0]))
}

func getFileString(file string) string {
	data, err := os.ReadFile(inputs + file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(data)
}

func solveOne(file string) int {
	data := getFileString(file)
	// split the file into lines
	lines := strings.Split(data, "\n")

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
	digitStrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

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

	// Find the int based digits and check if the first one is before the first text based digit,
	// and if the last one is after the last text based digit
	digit := regexp.MustCompile(`\d`)
	digits := digit.FindAllString(line, -1)

	if len(digits) >= 1 {
		if strings.Index(line, digits[0]) < firstIndex {
			firstDigit = digits[0]
		}

		if strings.Index(line, digits[len(digits)-1]) > lastIndex {
			lastDigit = digits[len(digits)-1]
		}
	}

	firstInt := convertDigitString(firstDigit)
	lastInt := convertDigitString(lastDigit)

	return firstInt*10 + lastInt
}

func solveTwo(file string) int {
	data := getFileString(file)

	// split the file into lines
	lines := strings.Split(data, "\n")

	var total int

	for _, line := range lines {
		total += findDigits(line)
	}
	return total
}
