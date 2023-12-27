package day5

import (
	"advent-of-code/common"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSeeds(t *testing.T) {
	assert.Equal(t, []int{79, 14, 55, 13}, getSeeds("inputs/test.txt"))
}

func TestGetSeedSoilMap(t *testing.T) {
	assert.Equal(t, [][]int{{50, 98, 2}, {52, 50, 48}}, getSeedToSoilMap("inputs/test.txt"))
}

func TestGetSoilFetilizerMap(t *testing.T) {
	assert.Equal(t, [][]int{{0, 15, 37}, {37, 52, 2}, {39, 0, 15}}, getSoilToFertilizerMap("inputs/test.txt"))
}

func TestGetFetilizerWaterMap(t *testing.T) {
	assert.Equal(t, [][]int{{49, 53, 8}, {0, 11, 42}, {42, 0, 7}, {57, 7, 4}}, getFertilizerToWaterMap("inputs/test.txt"))
}

func TestGetWaterLightMap(t *testing.T) {
	assert.Equal(t, [][]int{{88, 18, 7}, {18, 25, 70}}, getWaterToLightMap("inputs/test.txt"))
}

func TestGetLightTempMap(t *testing.T) {
	assert.Equal(t, [][]int{{45, 77, 23}, {81, 45, 19}, {68, 64, 13}}, getLightToTempMap("inputs/test.txt"))
}

func TestGetTempHumidityMap(t *testing.T) {
	assert.Equal(t, [][]int{{0, 69, 1}, {1, 0, 69}}, getTempToHumidityMap("inputs/test.txt"))
}

func TestGetHumidityLocMap(t *testing.T) {
	assert.Equal(t, [][]int{{60, 56, 37}, {56, 93, 4}}, getHumidityToLocMap("inputs/test.txt"))
}

func TestFindDestForSource(t *testing.T) {
	assert.Equal(t, 81, findDestForSource(79, []int{52, 50, 48}))
	assert.Equal(t, 79, findDestForSource(79, []int{50, 98, 2}))
}

func TestFindDestFromSources(t *testing.T) {
	assert.Equal(t, 81, findDestFromSources(79, [][]int{{52, 50, 48}, {50, 98, 2}}))
	assert.Equal(t, 14, findDestFromSources(14, [][]int{{52, 50, 48}, {50, 98, 2}}))
	assert.Equal(t, 57, findDestFromSources(55, [][]int{{52, 50, 48}, {50, 98, 2}}))
	assert.Equal(t, 13, findDestFromSources(13, [][]int{{52, 50, 48}, {50, 98, 2}}))
}

func TestFindLocationForSeed(t *testing.T) {
	m := [][][]int{{{50, 98, 2}, {52, 50, 48}},
		{{0, 15, 37}, {37, 52, 2}, {39, 0, 15}},
		{{49, 53, 8}, {0, 11, 42}, {42, 0, 7}, {57, 7, 4}},
		{{88, 18, 7}, {18, 25, 70}},
		{{45, 77, 23}, {81, 45, 19}, {68, 64, 13}},
		{{0, 69, 1}, {1, 0, 69}},
		{{60, 56, 37}, {56, 93, 4}},
	}
	assert.Equal(t, 82, findLocationForSeed(79, m))

}

func TestFindSeedFromLocation(t *testing.T) {
	m := [][][]int{{{50, 98, 2}, {52, 50, 48}},
		{{0, 15, 37}, {37, 52, 2}, {39, 0, 15}},
		{{49, 53, 8}, {0, 11, 42}, {42, 0, 7}, {57, 7, 4}},
		{{88, 18, 7}, {18, 25, 70}},
		{{45, 77, 23}, {81, 45, 19}, {68, 64, 13}},
		{{0, 69, 1}, {1, 0, 69}},
		{{60, 56, 37}, {56, 93, 4}},
	}

	slices.Reverse(m)
}

func TestFindLowestLocation(t *testing.T) {
	data := common.GetFileText("inputs/input.txt")
	seeds := getSeeds(data)
	soils := getSeedToSoilMap(data)
	fertilizer := getSoilToFertilizerMap(data)
	water := getFertilizerToWaterMap(data)
	light := getWaterToLightMap(data)
	temp := getLightToTempMap(data)
	humidity := getTempToHumidityMap(data)
	loc := getHumidityToLocMap(data)
	assert.Equal(t, 0, findLowestLocation(seeds, [][][]int{
		soils, fertilizer, water, light, temp, humidity, loc,
	}))
}

func TestSolveOne(t *testing.T) {
	assert.Equal(t, 13, SolveOne("inputs/test.txt"))
}

func TestSolveTwo(t *testing.T) {
	assert.Equal(t, 46, SolveTwo("inputs/test.txt"))
	assert.Equal(t, 63179500, SolveTwo("inputs/input.txt"))
}
