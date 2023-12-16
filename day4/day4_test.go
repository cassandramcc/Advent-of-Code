package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveOne(t *testing.T) {
	assert.Equal(t, 13, SolveOne("inputs/test.txt"))
}

func TestSolveTwo(t *testing.T) {
	assert.Equal(t, 30, SolveTwo("inputs/test.txt"))
}
