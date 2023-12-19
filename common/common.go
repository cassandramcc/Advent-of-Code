package common

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func GetFileLines(file string) []string {
	return strings.Split(strings.TrimSpace(GetFileText(file)), "\n")
}

func GetFileText(file string) string {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return string(data)
}

func ConvertToIntSlice(ss []string) []int {
	var result []int
	for _, s := range ss {
		i, _ := strconv.Atoi(s)
		result = append(result, i)
	}
	return result
}

func RemoveDuplicates(s []string) []string {
	inResult := make(map[string]bool)
	var result []string
	for _, str := range s {
		if _, ok := inResult[str]; !ok {
			inResult[str] = true
			result = append(result, str)
		}
	}
	return result
}

func CountInstances(ss []string) map[string]int {
	instancesMap := make(map[string]int)
	unduped := RemoveDuplicates(ss)
	for _, item := range ss {
		if slices.Contains(unduped, item) {
			instancesMap[item] += 1
		}
	}
	return instancesMap
}

func MapContains(m map[string]int, i int) bool {
	for _, v := range m {
		if v == i {
			return true
		}
	}
	return false
}

func MapFunc(ss []string, f func(s string) string) []string {
	var mappedSlice []string
	for _, s := range ss {
		mappedSlice = append(mappedSlice, f(s))
	}
	return mappedSlice
}
