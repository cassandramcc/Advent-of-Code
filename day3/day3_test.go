package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveOne(t *testing.T) {
	assert.Equal(t, 4361, SolveOne("inputs/test1.txt"))
}

func TestSolveTwo(t *testing.T) {
	assert.Equal(t, 467835, SolveTwo("inputs/test1.txt"))
}
