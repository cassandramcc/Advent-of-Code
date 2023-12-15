package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveOne(t *testing.T) {
	assert.Equal(t, 8, solveOne("inputs/test.txt"))
}

func TestSolveTwo(t *testing.T) {
	assert.Equal(t, 2286, solveTwo("inputs/test.txt"))
}
