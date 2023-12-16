package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveOne(t *testing.T) {
	assert.Equal(t, 4361, runOne("inputs/test1.txt"))
}

func TestSolveTwo(t *testing.T) {
	assert.Equal(t, 467835, runTwo("inputs/test1.txt"))
}
