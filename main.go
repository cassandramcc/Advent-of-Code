package main

import (
	"advent-of-code/day1"
	"advent-of-code/day2"
	"advent-of-code/day3"
	"advent-of-code/day4"
	"advent-of-code/day5"
	"advent-of-code/day6"
	"advent-of-code/day7"
	"advent-of-code/day8"
	"advent-of-code/day9"
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

	case "4":
		if part == "1" {
			fmt.Println(day4.SolveOne("day4/" + inputFile))
		} else if part == "2" {
			fmt.Println(day4.SolveTwo("day4/" + inputFile))
		}

	case "5":
		if part == "1" {
			fmt.Println(day5.SolveOne("day5/" + inputFile))
		} else if part == "2" {
			fmt.Println(day5.SolveTwo("day5/" + inputFile))
		}

	case "6":
		if part == "1" {
			fmt.Println(day6.SolveOne())
		} else if part == "2" {
			fmt.Println(day6.SolveTwo())
		}

	case "7":
		if part == "1" {
			fmt.Println(day7.SolveOne("day7/" + inputFile))
		} else if part == "2" {
			fmt.Println(day7.SolveTwo("day7/" + inputFile))
		}

	case "8":
		if part == "1" {
			fmt.Println(day8.SolveOne("day8/" + inputFile))
		} else if part == "2" {
			fmt.Println(day8.SolveTwo("day8/" + inputFile))
		}

	case "9":
		if part == "1" {
			fmt.Println(day9.SolveOne("day9/" + inputFile))
		} else if part == "2" {
			fmt.Println(day9.SolveTwo("day9/" + inputFile))
		}
	default:
		fmt.Println("No day for", day)
	}
}
