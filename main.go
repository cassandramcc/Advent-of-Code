package main

import (
	"advent-of-code/day1"
	"advent-of-code/day2"
	"advent-of-code/day3"
	"fmt"
	"os"
)

const (
	inputFile = "inputs/input.txt"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	if len(args) > 2 {
		fmt.Println("Too many arguments")
	}
	day := args[0]
	part := args[1]

	switch day {
	case "1":
		if part == "1" {
			fmt.Println(day1.SolveOne("day1/" + inputFile))
		} else if part == "2" {
			fmt.Println(day1.SolveTwo("day1/" + inputFile))
		}
	case "2":
		if part == "1" {
			fmt.Println(day2.SolveOne("day2/" + inputFile))
		} else if part == "2" {
			fmt.Println(day2.SolveTwo("day2/" + inputFile))
		}

	case "3":
		if part == "1" {
			fmt.Println(day3.SolveOne("day3/" + inputFile))
		} else if part == "2" {
			fmt.Println(day3.SolveTwo("day3/" + inputFile))
		}
	default:
		fmt.Println("No day for", day)
	}
}