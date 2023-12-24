package day9

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveOne(t *testing.T) {
	assert.Equal(t, 114, SolveOne("inputs/test.txt"))
}

func TestSolveTwo(t *testing.T) {
	assert.Equal(t, 2, SolveTwo("inputs/test.txt"))
}
