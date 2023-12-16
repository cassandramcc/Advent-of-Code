package day2

import (
	"advent-of-code/common"
	"regexp"
	"strconv"
)

func SolveOne(file string) int {
	lines := common.GetFileLines(file)

	var total int

	maxRed := 12
	maxGreen := 13
	maxBlue := 14

	for i, line := range lines {

		possible := true
		blueRegex := regexp.MustCompile(`(\d+) blue`)
		redRegex := regexp.MustCompile(`(\d+) red`)
		greenRegex := regexp.MustCompile(`(\d+) green`)

		blues := blueRegex.FindAllStringSubmatch(line, -1)
		reds := redRegex.FindAllStringSubmatch(line, -1)
		greens := greenRegex.FindAllStringSubmatch(line, -1)

		for _, blue := range blues {
			blueInt, _ := strconv.Atoi(blue[1])
			if blueInt > maxBlue {
				possible = false
				break
			}
		}

		for _, red := range reds {
			redInt, _ := strconv.Atoi(red[1])
			if redInt > maxRed {
				possible = false
				break
			}
		}

		for _, green := range greens {
			greenInt, _ := strconv.Atoi(green[1])
			if greenInt > maxGreen {
				possible = false
				break
			}
		}

		if possible {
			total = total + i + 1
		}
	}

	return total

}

func SolveTwo(file string) int {
	lines := common.GetFileLines(file)

	var total int

	for _, line := range lines {

		blueRegex := regexp.MustCompile(`(\d+) blue`)
		redRegex := regexp.MustCompile(`(\d+) red`)
		greenRegex := regexp.MustCompile(`(\d+) green`)

		blues := blueRegex.FindAllStringSubmatch(line, -1)
		reds := redRegex.FindAllStringSubmatch(line, -1)
		greens := greenRegex.FindAllStringSubmatch(line, -1)

		maxRed := 0
		maxGreen := 0
		maxBlue := 0

		for _, blue := range blues {
			blueInt, _ := strconv.Atoi(blue[1])
			if blueInt > maxBlue {
				maxBlue = blueInt
			}
		}

		for _, red := range reds {
			redInt, _ := strconv.Atoi(red[1])
			if redInt > maxRed {
				maxRed = redInt
			}
		}

		for _, green := range greens {
			greenInt, _ := strconv.Atoi(green[1])
			if greenInt > maxGreen {
				maxGreen = greenInt
			}
		}

		total = total + maxBlue*maxGreen*maxRed
	}

	return total
}
