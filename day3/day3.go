package day3

import (
	"advent-of-code/common"
	"regexp"
	"strconv"
)

func SolveTwo(file string) int {
	lines := common.GetFileLines(file)
	stars := findStars(lines)
	numbers := findNumbers(lines)
	gears := findGears(stars, numbers)
	return calculateTotal(gears, lines)
}

func SolveOne(file string) int {
	lines := common.GetFileLines(file)
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

		engineNumbers := findEngineNumbersOne(numbers, symbolsUp, symbolsDown, symbolsSame, line)
		for _, e := range engineNumbers {
			total += e
		}

	}

	return total
}

func findEngineNumbersOne(numbers, symbolsUp, symbolsDown, symbolsSame [][]int, line string) []int {
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

type number struct {
	x []int
	y int
}

func findStars(lines []string) [][]int {
	var stars [][]int
	starRegex := regexp.MustCompile(`\*`)
	for i, line := range lines {
		starIndexes := starRegex.FindAllStringIndex(line, -1)
		for _, starIndex := range starIndexes {
			stars = append(stars, []int{starIndex[0], i})
		}
	}
	return stars
}

func findNumbers(lines []string) []number {
	var numbers []number
	numbersRegex := regexp.MustCompile(`\d+`)
	for i, line := range lines {
		numberIndexes := numbersRegex.FindAllStringIndex(line, -1)
		for _, numberIndex := range numberIndexes {
			numbers = append(numbers, number{numberIndex, i})
		}
	}
	return numbers
}

func isAdjacent(star []int, num number) bool {
	return star[0] >= num.x[0]-1 && star[0] <= num.x[1] && (star[1] == num.y || star[1] == num.y+1 || star[1] == num.y-1)
}

func findGears(stars [][]int, numbers []number) [][]number {
	var gears [][]number

	for _, star := range stars {
		var adjacentNumbers []number
		for _, num := range numbers {
			if isAdjacent(star, num) {
				adjacentNumbers = append(adjacentNumbers, num)
			}
		}
		if len(adjacentNumbers) == 2 {
			gears = append(gears, adjacentNumbers)
		}
	}

	return gears
}

func calculateTotal(gears [][]number, lines []string) int {
	var total int

	for _, gear := range gears {
		gear1, _ := strconv.Atoi(lines[gear[0].y][gear[0].x[0]:gear[0].x[1]])
		gear2, _ := strconv.Atoi(lines[gear[1].y][gear[1].x[0]:gear[1].x[1]])
		total = total + (gear1 * gear2)
	}
	return total
}
