package day8

import (
	"advent-of-code/common"
	"regexp"
	"strings"
)

type nodes struct {
	left  string
	right string
}

func SolveOne(file string) int {
	lines := common.GetFileLines(file)
	network, instructions := getNetwork(lines)
	var steps int
	currentNode := "AAA"
	for currentNode != "ZZZ" {
		for _, i := range instructions {
			if currentNode == "ZZZ" {
				return steps
			}
			nextNodes := network[currentNode]
			nextNode := getNextLocation(nextNodes, i)
			steps += 1
			currentNode = nextNode
		}
	}
	return steps
}

func SolveTwo(file string) int {
	lines := common.GetFileLines(file)
	network, instructions := getNetwork(lines)
	currentNodes := findStartNodes(network)

	var recSteps []int
	for _, node := range currentNodes {
		recSteps = append(recSteps, followNodes(node, network, instructions))
	}

	return lcmm(recSteps)
}

func getNetwork(lines []string) (map[string]nodes, []string) {
	network := make(map[string]nodes)
	nodeRegex := regexp.MustCompile(`(\w{3})\s=\s\((\w{3}),\s(\w{3})\)`)
	for _, line := range lines[2:] {
		nodeMatches := nodeRegex.FindAllStringSubmatch(line, -1)
		network[nodeMatches[0][1]] = nodes{left: nodeMatches[0][2], right: nodeMatches[0][3]}
	}
	return network, strings.Split(lines[0], "")
}

func getNextLocation(nodes nodes, instruction string) string {
	if instruction == "L" {
		return nodes.left
	}
	return nodes.right
}

func findStartNodes(nodes map[string]nodes) []string {
	var startNodes []string
	for node := range nodes {
		if node[2] == 'A' {
			startNodes = append(startNodes, node)
		}
	}
	return startNodes
}

func isEndNode(currentNodes []string) bool {
	for _, node := range currentNodes {
		if node[2] != 'Z' {
			return false
		}
	}
	return true
}

func followNodes(startNode string, network map[string]nodes, instructions []string) int {
	var steps int
	currentNode := startNode
	for true {
		for _, i := range instructions {
			if isEndNode([]string{currentNode}) {
				return steps
			}
			nextNodes := network[currentNode]
			nextNode := getNextLocation(nextNodes, i)
			steps += 1
			currentNode = nextNode
		}
	}
	return -1
}

// gcd calculates the greatest common divisor using the Euclidean algorithm.
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// lcm calculates the lowest common multiple of two integers using the formula: lcm(a, b) = |a * b| / gcd(a, b)
func lcm(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	return (a * b) / gcd(a, b)
}

// lcmm calculates the lowest common multiple of multiple integers.
func lcmm(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}

	return result
}
