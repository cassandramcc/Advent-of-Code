package common

import (
	"fmt"
	"os"
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
