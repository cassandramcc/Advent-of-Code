package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	fmt.Println(runOne(args[0]))
}

func getFileString(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(data)
}

func runOne(file string) int {
	data := getFileString(file)

	return findTotal(data)

}

func findTotal(data string) int {
	lines := strings.Split(data, "\n")

	var total int

	numbersRegex := regexp.MustCompile(`\d+`)
	symbolsRegex := regexp.MustCompile(`[^\d^\.^\s]`)

	for i, line := range lines {
		var symbolsUp [][]int
		var symbolsDown [][]int
		var symbolsSame [][]int

		numbers := numbersRegex.FindAllStringIndex(line, -1)

		symbolsSame = symbolsRegex.FindAllStringIndex(line, -1)

		if i != 0 {
			symbolsUp = symbolsRegex.FindAllStringIndex(lines[i-1], -1)
		}

		if i != len(line)-1 {
			symbolsDown = symbolsRegex.FindAllStringIndex(lines[i+1], -1)
		}

		engineNumbers := findEngineNumbers(numbers, symbolsUp, symbolsDown, symbolsSame, line)
		for _, e := range engineNumbers {
			total += e
		}

	}

	return total
}

func findEngineNumbers(numbers, symbolsUp, symbolsDown, symbolsSame [][]int, line string) []int {
	var engineNumbers []int
	for _, num := range numbers {
		isEngineNumber := false
		for _, s := range symbolsUp {
			for idx := num[0]; idx < num[1]; idx++ {
				if idx == s[0] || idx == s[0]+1 || idx == s[0]-1 {
					isEngineNumber = true
				}
			}
		}

		for _, s := range symbolsDown {
			for idx := num[0]; idx < num[1]; idx++ {
				if idx == s[0] || idx == s[0]+1 || idx == s[0]-1 {
					isEngineNumber = true
				}
			}
		}

		for _, s := range symbolsSame {
			for idx := num[0]; idx < num[1]; idx++ {
				if idx == s[0] || idx == s[0]+1 || idx == s[0]-1 {
					isEngineNumber = true
				}
			}
		}

		if isEngineNumber {
			engineString := line[num[0]:num[1]]
			engineNumber, _ := strconv.Atoi(engineString)
			engineNumbers = append(engineNumbers, engineNumber)
		}
	}
	return engineNumbers
}
