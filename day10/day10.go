package day10

import (
	"advent-of-code/common"
	"fmt"
	"os"
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
	visitedPipes = append(visitedPipes, endPipe)
	nonLoop := findNonLoop(visitedPipes, grid)

	var count int
	for _, n := range nonLoop {
		if isInLoop(*n, visitedPipes, len(grid), len(grid[0])) {
			count++
			n.pipeType = "I"
		}
	}
	printGrid(visitedPipes, nonLoop, []int{len(grid), len(grid[0])})
	return count
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

func printGrid(pipes []pipe, nonPipe []*pipe, size []int) {
	colorRed := "\033[0;31m"
	colorNone := "\033[0m"
	colorBlue := "\033[0;34m"

	var grid [][]string
	for i := 0; i < size[0]; i++ {
		var row []string
		for j := 0; j < size[1]; j++ {
			row = append(row, ".")
		}
		grid = append(grid, row)
	}

	for _, p := range pipes {
		grid[p.y][p.x] = p.pipeType
	}

	for _, p := range nonPipe {
		grid[p.y][p.x] = p.pipeType
	}

	for y, g := range grid {
		for x, v := range g {
			if v == "I" {
				fmt.Fprintf(os.Stdout, "%s%s%s", colorRed, v, colorNone)
			} else if alreadyVisitedPipe(pipe{x: x, y: y}, pipes) {
				fmt.Fprintf(os.Stdout, "%s%s%s", colorBlue, v, colorNone)
			} else {
				fmt.Printf(v)
			}
		}
		fmt.Println()
	}
}

func findPipeAtLoc(p pipe, pipes []pipe) pipe {
	for _, v := range pipes {
		if v.x == p.x && v.y == p.y {
			return v
		}
	}
	return pipe{}
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

func findNonLoop(pipes []pipe, grid [][]string) []*pipe {
	var nonPipes []*pipe
	for i := range grid {
		for j := range grid[i] {
			p := pipe{pipeType: grid[i][j], x: j, y: i}
			if !alreadyVisitedPipe(p, pipes) {
				nonPipes = append(nonPipes, &p)
			}
		}
	}
	return nonPipes
}

func noStraights(m map[string]int) bool {
	_, ok := m["-"]
	return !ok
}

// This is the problem
// this is an attempt at the point in polygon problem with discrete ray casting
// however as it is discrete, it can't deal with when a point is inside the loop
// but on the vertical direction, there is no - pipe, only | type pipes.
// so if a point is inline with these vertical pipes, how to deal with it?

// if any of the counts are even then it's outside
func isInLoop(p pipe, loop []pipe, bottom, side int) bool {

	var verticalUpRay []string
	for y := p.y; y <= bottom; y++ {
		foundPipe := findPipeAtLoc(pipe{x: p.x, y: y}, loop)
		if alreadyVisitedPipe(pipe{x: p.x, y: y}, loop) {
			verticalUpRay = append(verticalUpRay, foundPipe.pipeType)
		}
	}

	instances := common.CountInstances(verticalUpRay)
	if !noStraights(instances) {
		return instances["-"]%2 != 0
	} else {
		total := 0
		for _, v := range instances {
			total += v
		}
		if total != 0 {
			return (total/2)%2 != 0
		}
		return false
	}

	var verticalDownRay []string
	for y := p.y; y >= 0; y-- {
		foundPipe := findPipeAtLoc(pipe{x: p.x, y: y}, loop)
		if alreadyVisitedPipe(pipe{x: p.x, y: y}, loop) {
			verticalDownRay = append(verticalDownRay, foundPipe.pipeType)
		}
	}

	instances = common.CountInstances(verticalDownRay)
	if !noStraights(instances) {
		return instances["-"]%2 != 0
	}

	var horizontalRightRay []string
	for x := p.x; x <= side; x++ {
		foundPipe := findPipeAtLoc(pipe{x: x, y: p.y}, loop)
		if alreadyVisitedPipe(pipe{x: x, y: p.y}, loop) {
			horizontalRightRay = append(horizontalRightRay, foundPipe.pipeType)
		}
	}

	instances = common.CountInstances(horizontalRightRay)
	if !noStraights(instances) {
		return instances["|"]%2 != 0
	}

	var horizontalLeftRay []string
	for x := p.x; x >= 0; x-- {
		foundPipe := findPipeAtLoc(pipe{x: x, y: p.y}, loop)
		if alreadyVisitedPipe(pipe{x: x, y: p.y}, loop) {
			horizontalLeftRay = append(horizontalLeftRay, foundPipe.pipeType)
		}
	}

	instances = common.CountInstances(horizontalLeftRay)
	if !noStraights(instances) {
		return instances["|"]%2 != 0
	}

	return false
}
