package day8

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveOneInputOne(t *testing.T) {
	assert.Equal(t, 2, SolveOne("inputs/test1.txt"))
}

func TestSolveOneInputTwo(t *testing.T) {
	assert.Equal(t, 6, SolveOne("inputs/test2.txt"))
}

func TestSolveTwoInputThree(t *testing.T) {
	assert.Equal(t, 6, SolveTwo("inputs/test3.txt"))
}
