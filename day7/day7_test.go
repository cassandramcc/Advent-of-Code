package day7

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolveOne(t *testing.T) {
	assert.Equal(t, 6440, SolveOne("inputs/test.txt"))
}
