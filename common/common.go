package common

import (
	"fmt"
	"os"
	"strings"
)

func GetFileLines(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(string(data), "\n")
}
