package day10

import (
	"advent-of-code/common"
	"strings"
)

type pipe struct {
	pipeType string
	x        int
	y        int
}

func SolveOne(file string) int {
	lines := common.GetFileLines(file)
	var grid [][]string
	sPipe := pipe{
		pipeType: "S",
		x:        0,
		y:        0,
	}
	for i, line := range lines {
		if strings.Contains(line, "S") {
			sPipe.x = strings.Index(line, "S")
			sPipe.y = i
		}
		grid = append(grid, strings.Split(line, ""))
	}

	var visitedPipes []pipe
	visitedPipes = append(visitedPipes, sPipe)
	pipes := findSPipeAdjacents(sPipe, grid)
	endPipe := pipes[1]
	currentPipe := pipes[0]
	for !isEndPipe(currentPipe, endPipe) {
		visitedPipes = append(visitedPipes, currentPipe)
		nextPipes := findNextPipe(currentPipe, grid)
		if !alreadyVisitedPipe(nextPipes[0], visitedPipes) {
			currentPipe = nextPipes[0]
		} else {
			currentPipe = nextPipes[1]
		}
	}
	return (len(visitedPipes) + 1) / 2
}

func SolveTwo(file string) int {
	return 0
}

func isEndPipe(p pipe, endPipe pipe) bool {
	return p.x == endPipe.x && p.y == endPipe.y
}
func alreadyVisitedPipe(p pipe, visted []pipe) bool {
	for _, v := range visted {
		if v.x == p.x && v.y == p.y {
			return true
		}
	}
	return false
}

func findNextPipe(p pipe, grid [][]string) []pipe {
	if p.pipeType == "7" {
		return []pipe{
			{
				pipeType: grid[p.y+1][p.x],
				x:        p.x,
				y:        p.y + 1,
			},
			{
				pipeType: grid[p.y][p.x-1],
				x:        p.x - 1,
				y:        p.y,
			},
		}
	} else if p.pipeType == "-" {
		return []pipe{
			{
				pipeType: grid[p.y][p.x-1],
				x:        p.x - 1,
				y:        p.y,
			},
			{
				pipeType: grid[p.y][p.x+1],
				x:        p.x + 1,
				y:        p.y,
			},
		}
	} else if p.pipeType == "L" {
		return []pipe{
			{
				pipeType: grid[p.y][p.x+1],
				x:        p.x + 1,
				y:        p.y,
			},
			{
				pipeType: grid[p.y-1][p.x],
				x:        p.x,
				y:        p.y - 1,
			},
		}
	} else if p.pipeType == "J" {
		return []pipe{
			{
				pipeType: grid[p.y][p.x-1],
				x:        p.x - 1,
				y:        p.y,
			},
			{
				pipeType: grid[p.y-1][p.x],
				x:        p.x,
				y:        p.y - 1,
			},
		}
	} else if p.pipeType == "F" {
		return []pipe{
			{
				pipeType: grid[p.y][p.x+1],
				x:        p.x + 1,
				y:        p.y,
			},
			{
				pipeType: grid[p.y+1][p.x],
				x:        p.x,
				y:        p.y + 1,
			},
		}
	} else if p.pipeType == "|" {
		return []pipe{
			{
				pipeType: grid[p.y-1][p.x],
				x:        p.x,
				y:        p.y - 1,
			},
			{
				pipeType: grid[p.y+1][p.x],
				x:        p.x,
				y:        p.y + 1,
			},
		}
	}
	return nil
}

func findSPipeAdjacents(sPipe pipe, grid [][]string) []pipe {
	// right and down
	if sPipe.x+1 <= len(grid[0]) && sPipe.y+1 < len(grid) && (grid[sPipe.y][sPipe.x+1] == "7" || grid[sPipe.y][sPipe.x+1] == "-" || grid[sPipe.y][sPipe.x+1] == "J") && (grid[sPipe.y+1][sPipe.x] == "|" || grid[sPipe.y+1][sPipe.x] == "L" || grid[sPipe.y+1][sPipe.x] == "J") {

		pipes := []pipe{{
			pipeType: grid[sPipe.y][sPipe.x+1],
			x:        sPipe.x + 1,
			y:        sPipe.y,
		}, {
			pipeType: grid[sPipe.y+1][sPipe.x],
			x:        sPipe.x,
			y:        sPipe.y + 1,
		}}
		return pipes
	}

	if sPipe.y+1 <= len(grid) && sPipe.x-1 >= 0 && (grid[sPipe.y+1][sPipe.x] == "|" || grid[sPipe.y+1][sPipe.x] == "L" || grid[sPipe.y+1][sPipe.x] == "J") && (grid[sPipe.y][sPipe.x-1] == "-" || grid[sPipe.y][sPipe.x-1] == "L" || grid[sPipe.y][sPipe.x-1] == "F") {
		pipes := []pipe{{
			pipeType: grid[sPipe.y+1][sPipe.x],
			x:        sPipe.x,
			y:        sPipe.y + 1,
		},
			{
				pipeType: grid[sPipe.y][sPipe.x-1],
				x:        sPipe.x - 1,
				y:        sPipe.y,
			}}
		return pipes
	}

	if sPipe.x-1 >= 0 && sPipe.y-1 >= 0 && (grid[sPipe.y][sPipe.x-1] == "-" || grid[sPipe.y][sPipe.x-1] == "L" || grid[sPipe.y][sPipe.x-1] == "F") && (grid[sPipe.y-1][sPipe.x] == "|" || grid[sPipe.y-1][sPipe.x] == "F" || grid[sPipe.y-1][sPipe.x] == "7") {
		pipes := []pipe{{
			pipeType: grid[sPipe.y][sPipe.x-1],
			x:        sPipe.x - 1,
			y:        sPipe.y,
		}, {
			pipeType: grid[sPipe.y-1][sPipe.x],
			x:        sPipe.x,
			y:        sPipe.y - 1,
		}}
		return pipes
	}

	if sPipe.y-1 >= 0 && sPipe.x+1 <= len(grid[0]) && (grid[sPipe.y-1][sPipe.x] == "|" || grid[sPipe.y-1][sPipe.x] == "F" || grid[sPipe.y-1][sPipe.x] == "7") && (grid[sPipe.y][sPipe.x+1] == "-" || grid[sPipe.y][sPipe.x+1] == "J" || grid[sPipe.y][sPipe.x+1] == "7") {
		pipes := []pipe{{
			pipeType: grid[sPipe.y-1][sPipe.x],
			x:        sPipe.x,
			y:        sPipe.y - 1,
		}, {
			pipeType: grid[sPipe.y][sPipe.x+1],
			x:        sPipe.x + 1,
			y:        sPipe.y,
		}}
		return pipes
	}
	return nil
}
