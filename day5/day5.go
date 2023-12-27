package day5

import (
	"advent-of-code/common"
	"fmt"
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

func SolveTwo(file string) int {
	data := common.GetFileText(file)
	seedRanges := getSeedRanges(data)
	maps := getMaps(file) // soils, fertilizer, water, light, temp, humidity, location

	var possibleLocations []int
	for _, soils := range maps[1] {
		possibleLocations = append(possibleLocations, findLocationForSeed(soils[0], maps))
		fmt.Println("seed->", soils[0], "location", findLocationForSeed(soils[0], maps))
		possibleLocations = append(possibleLocations, findLocationForSeed(soils[1], maps[1:]))
		fmt.Println("soil->", soils[1], "location", findLocationForSeed(soils[1], maps[1:]))
	}

	for _, fertilizer := range maps[2] {
		possibleLocations = append(possibleLocations, findLocationForSeed(fertilizer[1], maps[2:]))
		fmt.Println("fertilizer->", fertilizer[1], "location", findLocationForSeed(fertilizer[1], maps[2:]))
	}

	for _, water := range maps[3] {
		possibleLocations = append(possibleLocations, findLocationForSeed(water[1], maps[3:]))
		fmt.Println("water->", water[1], "location", findLocationForSeed(water[1], maps[3:]))
	}

	for _, light := range maps[4] {
		possibleLocations = append(possibleLocations, findLocationForSeed(light[1], maps[4:]))
		fmt.Println("light->", light[1], "location", findLocationForSeed(light[1], maps[4:]))
	}

	for _, temp := range maps[5] {
		possibleLocations = append(possibleLocations, findLocationForSeed(temp[1], maps[5:]))
		fmt.Println("temp->", temp[1], "location", findLocationForSeed(temp[1], maps[5:]))
	}

	for _, humid := range maps[6] {
		possibleLocations = append(possibleLocations, findLocationForSeed(humid[1], maps[6:]))
		fmt.Println("humidity->", humid[1], "location", findLocationForSeed(humid[1], maps[6:]))
	}

	for _, loc := range maps[6] {
		possibleLocations = append(possibleLocations, loc[1])
		fmt.Println("loc->", loc[1], "location", loc[1])
	}

	// loop through each location in order
	slices.Sort(possibleLocations)
	fmt.Println(possibleLocations)

	// reverse the maps so a seed can be found for a location
	slices.Reverse(maps)
	for _, l := range possibleLocations {
		seed := findSeedFromLocation(l, maps)
		fmt.Println("location->", l, "seed", seed)
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
