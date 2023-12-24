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
