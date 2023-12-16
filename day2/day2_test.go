package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveOne(t *testing.T) {
	assert.Equal(t, 8, SolveOne("inputs/test.txt"))
}

func TestSolveTwo(t *testing.T) {
	assert.Equal(t, 2286, SolveTwo("inputs/test.txt"))
}
