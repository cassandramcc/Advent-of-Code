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
	soils := getSeedToSoilMap(data)
	fertilizer := getSoilToFertilizerMap(data)
	water := getFertilizerToWaterMap(data)
	light := getWaterToLightMap(data)
	temp := getLightToTempMap(data)
	humidity := getTempToHumidityMap(data)
	loc := getHumidityToLocMap(data)

	return findLowestLocation(seeds, [][][]int{
		soils, fertilizer, water, light, temp, humidity, loc,
	})
}

func SolveTwo(file string) int {
	return 0
}

func getSeeds(data string) []int {
	seedsRegex := regexp.MustCompile(`seeds:((\s+\d+)+)`)
	digitsRegex := regexp.MustCompile(`\d+`)
	seedsMatches := seedsRegex.FindStringSubmatch(data)
	digitStrings := digitsRegex.FindAllString(seedsMatches[1], -1)
	return common.ConvertToIntSlice(digitStrings)
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
