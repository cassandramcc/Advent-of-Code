package day10

import (
	"advent-of-code/common"
	"fmt"
	"strings"
)

type pipe struct {
	pipeType string
	x        int
	y        int
	input    *pipe
	output   *pipe
}

func SolveOne(file string) int {
	lines := common.GetFileLines(file)
	var grid [][]string
	sPipe := pipe{
		pipeType: "",
		x:        0,
		y:        0,
		input:    nil,
		output:   nil,
	}
	for i, line := range lines {
		if strings.Contains(line, "S") {
			sPipe.x = strings.Index(line, "S")
			sPipe.y = i
		}
		grid = append(grid, strings.Split(line, ""))
	}

	pipes := findSPipeAdjacents(sPipe, grid)
	fmt.Println(pipes)

	return 0
}

func SolveTwo(file string) int {
	//data := common.GetFileText(file)
	return 0
}

func isValidPipeChar(s string) bool {
	return s != "."
}

func findSPipeAdjacents(sPipe pipe, grid [][]string) []pipe {
	if sPipe.x+1 <= len(grid[0]) && sPipe.y+1 < len(grid) && (grid[sPipe.y][sPipe.x+1] == "7" || grid[sPipe.y][sPipe.x+1] == "-" || grid[sPipe.y][sPipe.x+1] == "J") && (grid[sPipe.y+1][sPipe.x] == "|" || grid[sPipe.y+1][sPipe.x] == "L" || grid[sPipe.y+1][sPipe.x] == "J") {

		pipes := []pipe{{
			pipeType: grid[sPipe.y][sPipe.x+1],
			x:        sPipe.x + 1,
			y:        sPipe.y,
			input:    &sPipe,
			output:   nil,
		}, {
			pipeType: grid[sPipe.y+1][sPipe.x],
			x:        sPipe.x,
			y:        sPipe.y,
			input:    nil,
			output:   &sPipe,
		}}
		sPipe.output = &pipes[0]
		sPipe.input = &pipes[1]
		return pipes
	}

	if sPipe.y+1 <= len(grid) && sPipe.x-1 >= 0 && (grid[sPipe.y+1][sPipe.x] == "|" || grid[sPipe.y+1][sPipe.x] == "L" || grid[sPipe.y+1][sPipe.x] == "J") && (grid[sPipe.y][sPipe.x-1] == "-" || grid[sPipe.y][sPipe.x-1] == "L" || grid[sPipe.y][sPipe.x-1] == "F") {
		pipes := []pipe{{
			pipeType: grid[sPipe.y+1][sPipe.x],
			x:        sPipe.x,
			y:        sPipe.y - 1,
			input:    &sPipe,
			output:   nil,
		},
			{
				pipeType: grid[sPipe.y][sPipe.x-1],
				x:        sPipe.x - 1,
				y:        sPipe.y,
				input:    nil,
				output:   &sPipe,
			}}
		sPipe.output = &pipes[0]
		sPipe.input = &pipes[1]
		return pipes
	}

	if sPipe.x-1 >= 0 && sPipe.y-1 >= 0 && (grid[sPipe.y][sPipe.x-1] == "-" || grid[sPipe.y][sPipe.x-1] == "L" || grid[sPipe.y][sPipe.x-1] == "F") && (grid[sPipe.y-1][sPipe.x] == "|" || grid[sPipe.y-1][sPipe.x] == "F" || grid[sPipe.y-1][sPipe.x] == "7") {
		pipes := []pipe{{
			pipeType: grid[sPipe.y][sPipe.x-1],
			x:        sPipe.x - 1,
			y:        sPipe.y,
			input:    &sPipe,
			output:   nil,
		}, {
			pipeType: grid[sPipe.y-1][sPipe.x],
			x:        sPipe.x,
			y:        sPipe.y + 1,
			input:    nil,
			output:   &sPipe,
		}}
		sPipe.output = &pipes[0]
		sPipe.input = &pipes[1]
		return pipes
	}

	if sPipe.y-1 >= 0 && sPipe.x+1 <= len(grid[0]) && (grid[sPipe.y-1][sPipe.x] == "|" || grid[sPipe.y-1][sPipe.x] == "F" || grid[sPipe.y-1][sPipe.x] == "7") && (grid[sPipe.y][sPipe.x+1] == "-" || grid[sPipe.y][sPipe.x+1] == "J" || grid[sPipe.y][sPipe.x+1] == "7") {
		pipes := []pipe{{
			pipeType: grid[sPipe.y-1][sPipe.x],
			x:        sPipe.x,
			y:        sPipe.y + 1,
			input:    &sPipe,
			output:   nil,
		}, {
			pipeType: grid[sPipe.y][sPipe.x+1],
			x:        sPipe.x + 1,
			y:        sPipe.y,
			input:    nil,
			output:   &sPipe,
		}}
		sPipe.output = &pipes[0]
		sPipe.input = &pipes[1]
		return pipes
	}
	return nil
}
