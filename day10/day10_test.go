package day10

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveOneTestOne(t *testing.T) {
	assert.Equal(t, 4, SolveOne("inputs/test1.txt"))
}

func TestSolveOneTestTwo(t *testing.T) {
	assert.Equal(t, 8, SolveOne("inputs/test2.txt"))
}

func TestSolveTwoTestOne(t *testing.T) {
	assert.Equal(t, 1, SolveTwo("inputs/test1.txt"))
}

func TestSolveTwoTestThree(t *testing.T) {
	assert.Equal(t, 4, SolveTwo("inputs/test3.txt"))
}

func TestSolveTwoTestFour(t *testing.T) {
	assert.Equal(t, 4, SolveTwo("inputs/test4.txt"))
}

func TestSolveTwoTestFive(t *testing.T) {
	assert.Equal(t, 8, SolveTwo("inputs/test5.txt"))
}

func TestSolveTwoTestSix(t *testing.T) {
	assert.Equal(t, 10, SolveTwo("inputs/test6.txt"))
}

func TestSolveTwoTestSeven(t *testing.T) {
	assert.Equal(t, 1, SolveTwo("inputs/test7.txt"))
}
