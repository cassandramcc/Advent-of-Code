package day5

import (
	"advent-of-code/common"
	"regexp"
	"slices"
	"strings"
)

func SolveOne(file string) int {
	data := common.GetFileText(file)
	seeds := getSeeds(data)
	maps := getMaps(file)

	return findLowestLocation(seeds, maps)
}

func findPossibleLocations(ranges [][]int, s string, maps [][][]int) []int {
	var possibleLocations []int
	for _, r := range ranges {
		// r[0] is the GODDAMN DESTINATION
		// i.e. in a seed to soil map, r[0] IS THE GODDAMN SOIL
		possibleLocations = append(possibleLocations, findLocationForSeed(r[0], maps))
	}

	return possibleLocations
}

func SolveTwo(file string) int {
	data := common.GetFileText(file)
	seedRanges := getSeedRanges(data)
	maps := getMaps(file) // seed to soil, soil to fertilizer, fertilizer to water, water to light, light to temp, temp to humidity, humidity to location

	var possibleLocations []int
	possibleLocations = append(possibleLocations, findPossibleLocations(maps[0], "soil", maps[1:])...)
	possibleLocations = append(possibleLocations, findPossibleLocations(maps[1], "fertilizer", maps[2:])...)
	possibleLocations = append(possibleLocations, findPossibleLocations(maps[2], "water", maps[3:])...)
	possibleLocations = append(possibleLocations, findPossibleLocations(maps[3], "light", maps[4:])...)
	possibleLocations = append(possibleLocations, findPossibleLocations(maps[4], "temp", maps[5:])...)
	possibleLocations = append(possibleLocations, findPossibleLocations(maps[5], "humidity", maps[6:])...)

	// loop through each location in order
	slices.Sort(possibleLocations)

	// reverse the maps so a seed can be found for a location
	slices.Reverse(maps)
	for _, l := range possibleLocations {
		seed := findSeedFromLocation(l, maps)
		if isValidSeed(seed, seedRanges) {
			return l
		}
	}

	return 0
}

func getMaps(file string) [][][]int {
	data := common.GetFileText(file)
	soils := getSeedToSoilMap(data)
	fertilizer := getSoilToFertilizerMap(data)
	water := getFertilizerToWaterMap(data)
	light := getWaterToLightMap(data)
	temp := getLightToTempMap(data)
	humidity := getTempToHumidityMap(data)
	loc := getHumidityToLocMap(data)
	return [][][]int{soils, fertilizer, water, light, temp, humidity, loc}
}

func getSeeds(data string) []int {
	seedsRegex := regexp.MustCompile(`seeds:((\s+\d+)+)`)
	digitsRegex := regexp.MustCompile(`\d+`)
	seedsMatches := seedsRegex.FindStringSubmatch(data)
	digitStrings := digitsRegex.FindAllString(seedsMatches[1], -1)
	return common.ConvertToIntSlice(digitStrings)
}

func getSeedRanges(data string) [][]int {
	var seedRanges [][]int
	seeds := getSeeds(data)
	for i := 0; i < len(seeds); i += 2 {
		seedRanges = append(seedRanges, []int{seeds[i], seeds[i+1]})
	}
	return seedRanges
}

func getSeedToSoilMap(data string) [][]int {
	seedSoilRegex := regexp.MustCompile(`seed-to-soil map:(\s+(?:\d+\s)+)+`)
	digitsRegex := regexp.MustCompile(`\d+`)
	seedSoilMatch := seedSoilRegex.FindStringSubmatch(data)
	var seedSoilMap [][]int
	rows := strings.Split(seedSoilMatch[1], "\n")
	for _, row := range rows {
		digitsMatch := digitsRegex.FindAllString(row, -1)
		if digitsMatch != nil {
			seedSoilMap = append(seedSoilMap, common.ConvertToIntSlice(digitsMatch))
		}
	}
	return seedSoilMap
}

func getSoilToFertilizerMap(data string) [][]int {
	mapRegex := regexp.MustCompile(`soil-to-fertilizer map:(\s+(?:\d+\s)+)+`)
	digitsRegex := regexp.MustCompile(`\d+`)
	mapMatch := mapRegex.FindStringSubmatch(data)
	var intMap [][]int
	rows := strings.Split(mapMatch[1], "\n")
	for _, row := range rows {
		digitsMatch := digitsRegex.FindAllString(row, -1)
		if digitsMatch != nil {
			intMap = append(intMap, common.ConvertToIntSlice(digitsMatch))
		}
	}
	return intMap
}

func getFertilizerToWaterMap(data string) [][]int {
	mapRegex := regexp.MustCompile(`fertilizer-to-water map:(\s+(?:\d+\s)+)+`)
	digitsRegex := regexp.MustCompile(`\d+`)
	mapMatch := mapRegex.FindStringSubmatch(data)
	var intMap [][]int
	rows := strings.Split(mapMatch[1], "\n")
	for _, row := range rows {
		digitsMatch := digitsRegex.FindAllString(row, -1)
		if digitsMatch != nil {
			intMap = append(intMap, common.ConvertToIntSlice(digitsMatch))
		}
	}
	return intMap
}

func getWaterToLightMap(data string) [][]int {
	mapRegex := regexp.MustCompile(`water-to-light map:(\s+(?:\d+\s)+)+`)
	digitsRegex := regexp.MustCompile(`\d+`)
	mapMatch := mapRegex.FindStringSubmatch(data)
	var intMap [][]int
	rows := strings.Split(mapMatch[1], "\n")
	for _, row := range rows {
		digitsMatch := digitsRegex.FindAllString(row, -1)
		if digitsMatch != nil {
			intMap = append(intMap, common.ConvertToIntSlice(digitsMatch))
		}
	}
	return intMap
}

func getLightToTempMap(data string) [][]int {
	mapRegex := regexp.MustCompile(`light-to-temperature map:(\s+(?:\d+\s)+)+`)
	digitsRegex := regexp.MustCompile(`\d+`)
	mapMatch := mapRegex.FindStringSubmatch(data)
	var intMap [][]int
	rows := strings.Split(mapMatch[1], "\n")
	for _, row := range rows {
		digitsMatch := digitsRegex.FindAllString(row, -1)
		if digitsMatch != nil {
			intMap = append(intMap, common.ConvertToIntSlice(digitsMatch))
		}
	}
	return intMap
}

func getTempToHumidityMap(data string) [][]int {
	mapRegex := regexp.MustCompile(`temperature-to-humidity map:(\s+(?:\d+\s)+)+`)
	digitsRegex := regexp.MustCompile(`\d+`)
	mapMatch := mapRegex.FindStringSubmatch(data)
	var intMap [][]int
	rows := strings.Split(mapMatch[1], "\n")
	for _, row := range rows {
		digitsMatch := digitsRegex.FindAllString(row, -1)
		if digitsMatch != nil {
			intMap = append(intMap, common.ConvertToIntSlice(digitsMatch))
		}
	}
	return intMap
}

func getHumidityToLocMap(data string) [][]int {
	mapRegex := regexp.MustCompile(`humidity-to-location map:(\s+(?:\d+\s*)+)+`)
	digitsRegex := regexp.MustCompile(`\d+`)
	mapMatch := mapRegex.FindStringSubmatch(data)
	var intMap [][]int
	rows := strings.Split(mapMatch[1], "\n")
	for _, row := range rows {
		digitsMatch := digitsRegex.FindAllString(row, -1)
		if digitsMatch != nil {
			intMap = append(intMap, common.ConvertToIntSlice(digitsMatch))
		}
	}
	return intMap
}

// Find the destination of a source from one range information line
func findDestForSource(source int, ranges []int) int {
	if ranges[1] <= source && source < ranges[1]+ranges[2] {
		return (source - ranges[1]) + ranges[0]
	}
	return source
}

func isValidSeed(seed int, seedRanges [][]int) bool {
	for _, r := range seedRanges {
		if isInRange(seed, r) {
			return true
		}
	}
	return false
}

func isInRange(source int, ranges []int) bool {
	return ranges[0] <= source && source < ranges[0]+ranges[1]
}

// Find the source of a destination from one range information line
func findSourceForDest(dest int, ranges []int) int {
	if ranges[0] <= dest && dest < ranges[0]+ranges[2] {
		return (dest - ranges[0]) + ranges[1]
	}
	return dest
}

func findSourceFromDests(dest int, ranges [][]int) int {
	for _, r := range ranges {
		source := findSourceForDest(dest, r)
		if source != dest {
			return source
		}
	}
	return dest
}

func findSeedFromLocation(location int, maps [][][]int) int {
	source := location
	for _, m := range maps {
		source = findSourceFromDests(source, m)
	}
	return source
}

func findDestFromSources(source int, ranges [][]int) int {
	for _, r := range ranges {
		dest := findDestForSource(source, r)
		if dest != source {
			return dest
		}
	}
	return source
}

func findLocationForSeed(seed int, maps [][][]int) int {
	dest := seed
	for _, m := range maps {
		dest = findDestFromSources(dest, m)
		//fmt.Println(dest)
	}
	return dest
}

func findLowestLocation(seeds []int, maps [][][]int) int {
	var locs []int
	for _, seed := range seeds {
		locs = append(locs, findLocationForSeed(seed, maps))
	}
	return slices.Min(locs)
}
