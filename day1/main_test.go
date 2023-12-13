package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFileString(t *testing.T) {
	assert.Equal(t, "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet", getFileString("test1.txt"))
}

func TestSolveOne(t *testing.T) {
	assert.Equal(t, 142, solveOne("test1.txt"))
}

func TestSolveTwo(t *testing.T) {
	assert.Equal(t, 281, solveTwo("test2.txt"))
}

func TestConvertDigitString(t *testing.T) {
	tt := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	for tc, exp := range tt {
		t.Run(tc, func(t *testing.T) {
			assert.Equal(t, exp, convertDigitString(tc))
		})
	}
}

func TestFindDigits(t *testing.T) {
	tt := map[string]int{
		"two1nine":                               29,
		"4nineeightseven2":                       42,
		"fivethreeonezblqnsfk1":                  51,
		"four156":                                46,
		"3qcf":                                   33,
		"6gd9":                                   69,
		"z5":                                     55,
		"3five5":                                 35,
		"499jvmmrfzkmbppzcm":                     49,
		"66":                                     66,
		"four1dcczj":                             41,
		"1rdtwofjvdllht5eightsixfourbl":          14,
		"3four9":                                 39,
		"zgnoneightseveneightseven5d2fivefourjp": 14,
		"fouronetwo7seventwoeight":               48,
		"2lssdgdvhl":                             22,
		"zjlgvvpzpone9":                          19,
		"doneightn":                              18,
		"zoneight234":                            14,
		"oneightone":                             11,
	}

	for tc, exp := range tt {
		t.Run(tc, func(t *testing.T) {
			assert.Equal(t, exp, findDigits(tc))
		})
	}
}

func TestOverlaps(t *testing.T) {
	assert.Equal(t, 99, findDigits("nine7777nine"))
}
