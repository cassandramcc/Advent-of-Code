package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	inputs = "inputs/"
	test   = "test.txt"
	part1  = "part1.txt"
)

func main() {
	solve(test)
}

func getFileString(file string) string {
	data, err := os.ReadFile(inputs + file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(data)
}

func solve(file string) int {
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
